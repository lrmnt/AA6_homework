package cron

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/billing/ent"
	"github.com/lrmnt/AA6_homework/billing/ent/billing"
	"github.com/lrmnt/AA6_homework/billing/ent/billingcycle"
	"github.com/lrmnt/AA6_homework/billing/ent/operations"
	"github.com/lrmnt/AA6_homework/billing/ent/user"
	"github.com/lrmnt/AA6_homework/lib/api/schema"
	"github.com/lrmnt/AA6_homework/lib/api/schema/billing_event"
	"github.com/lrmnt/AA6_homework/lib/kafka"
	"go.uber.org/zap"
	"math"
	"time"
)

type Service struct {
	ticker   *time.Ticker
	client   *ent.Client
	log      *zap.Logger
	producer *kafka.Producer
}

func New(log *zap.Logger, client *ent.Client, producer *kafka.Producer) *Service {
	return &Service{
		ticker:   time.NewTicker(time.Minute),
		log:      log,
		client:   client,
		producer: producer,
	}
}

func (s *Service) tx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("can not start tx: %w", err)
	}

	err = fn(tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("can not tollback tx: %w, Err: %w", rbErr, err)
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("can not commit tx: %w", err)
	}

	return nil
}

func (s *Service) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			s.ticker.Stop()
			return ctx.Err()
		case <-s.ticker.C:
			err := s.checkAndMakePayout(ctx)
			if err != nil {
				s.log.Error("can not make payout", zap.Error(err))
			}
		}
	}
}

func (s *Service) checkAndMakePayout(ctx context.Context) error {

	bc, err := s.getOneEndedBC(ctx)
	if err != nil {
		if ent.IsNotFound(err) { // no opened billing cycles to process
			return nil
		}

		return err
	}

	s.log.Info("starting payout process", zap.Int("bc_id", bc.ID))

	// get all users
	users, err := s.client.User.Query().All(ctx)
	if err != nil {
		return err
	}

	for i := range users {
		err = s.processOneUserBilling(ctx, users[i], bc)
		if err != nil {
			return err
		}
	}

	// close billing cycle
	err = s.client.BillingCycle.UpdateOne(bc).
		SetState(billingcycle.StateClosed).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) processOneUserBilling(ctx context.Context, u *ent.User, bc *ent.BillingCycle) error {
	err := s.tx(ctx, func(tx *ent.Tx) error {
		// check that user have no billing operation in that billing cycle
		exists, err := tx.Billing.Query().
			Where(billing.And(
				billing.HasUserWith(user.ID(u.ID)),
				billing.HasBillingCycleWith(billingcycle.ID(bc.ID)),
			)).
			Exist(ctx)
		if err != nil {
			return err
		}
		if exists {
			return nil // already paid
		}

		// get user operations in current billing cycle
		userOps, err := tx.Operations.Query().
			Where(operations.And(
				operations.HasUserWith(user.ID(u.ID)),
				operations.HasBillingCycleWith(billingcycle.ID(bc.ID)),
			)).
			All(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		if len(userOps) == 0 {
			return nil
		}

		// calculate daily amount
		var bcdiff int64
		for j := range userOps {
			switch userOps[j].Type {
			case operations.TypeIncome:
				bcdiff += userOps[j].Amount
			case operations.TypeOutcome:
				bcdiff -= userOps[j].Amount
			default:
				s.log.Error("unknown operation type", zap.String("operation_type", userOps[j].Type.String()))
			}
		}

		newBalance := u.Balance + bcdiff

		var (
			billingOperationType billing.Type
			billingEventType     billing_event.Event
		)
		if newBalance > 0 {
			billingOperationType = billing.TypePayout
			billingEventType = billing_event.Event_EVENT_PAYED_TO_USER

			err = u.Update().SetBalance(0).Exec(ctx)
			if err != nil {
				return err
			}
		} else {
			billingOperationType = billing.TypeDebtMove
			billingEventType = billing_event.Event_EVENT_USER_BALANCE_BELOW_ZERO_AT_THE_END_OF_DAY

			err = u.Update().SetBalance(newBalance).Exec(ctx)
			if err != nil {
				return err
			}
		}

		// create payout or balance move in DB
		err = tx.Billing.Create().
			SetBillingCycleID(bc.ID).
			SetUUID(uuid.New()).
			SetUserID(u.ID).
			SetAmount(newBalance).
			SetTimestamp(time.Now().UnixNano()).
			SetType(billingOperationType).
			Exec(ctx)
		if err != nil {
			return err
		}

		if billingOperationType == billing.TypePayout {
			// pay money, in real world there could be an error from provider, but we are still able to rollback tx
			s.log.Warn("payed for user",
				zap.String("user public id", u.UUID.String()),
				zap.Int64("amount", newBalance),
				zap.Int64("bc start", bc.TsFrom),
				zap.Int64("bc end", bc.TsTo))
		}

		// send event to kafka

		mes := &billing_event.BillingEventV1{
			Event:     billingEventType,
			Timestamp: time.Now().UnixNano(),
			EventId:   uuid.New().String(),
			UserId:    u.UUID.String(),
			Amount:    int64(math.Abs(float64(newBalance))),
		}

		_, err = schema.ValidateBillingEventV1(mes)
		if err != nil {
			return err
		}

		data, err := proto.Marshal(mes)
		if err != nil {
			return err
		}

		return s.producer.Produce(data)
	})

	return err
}

func (s *Service) getOneEndedBC(ctx context.Context) (*ent.BillingCycle, error) {
	ts := time.Now().UnixNano()

	return s.client.BillingCycle.Query().
		Where(billingcycle.And(
			billingcycle.TsToLT(ts), // less than
			billingcycle.StateIn(billingcycle.StateOpen),
		)).
		Limit(1).
		Only(ctx)
}
