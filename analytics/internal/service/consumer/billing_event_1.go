package consumer

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/analytics/ent/billingoperations"
	"github.com/lrmnt/AA6_homework/analytics/ent/user"
	"github.com/lrmnt/AA6_homework/lib/api/schema"
	"github.com/lrmnt/AA6_homework/lib/api/schema/billing_event"
	"go.uber.org/zap"
)

func (s *Service) RunConsumeBillingEventV1(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			mes, err := s.billingEventV1Consumer.FetchMessage(ctx)
			if err != nil {
				s.log.Error("can not read message from queue", zap.Error(err))
				continue
			}

			var billingEvent billing_event.BillingEventV1

			err = proto.Unmarshal(mes.Value, &billingEvent)
			if err != nil {
				s.log.Error("can not unmarshal billing event from queue", zap.Error(err))
				continue
			}

			ok, err := schema.ValidateBillingEventV1(&billingEvent)
			if err != nil || !ok { // invalid message format
				s.log.Warn("can not validate billing event v1 message")

				err = s.billingEventV1Consumer.CommitMessages(ctx, mes)
				if err != nil {
					s.log.Error("can not commit billing event", zap.Error(err))
				}

				continue
			}

			err = s.processBillingEventV1(ctx, &billingEvent)
			if err != nil {
				s.log.Error("can not process billing event", zap.Error(err))
				continue
			}

			err = s.billingEventV1Consumer.CommitMessages(ctx, mes)
			if err != nil {
				s.log.Error("can not commit billing event", zap.Error(err))
			}

		}
	}
}

func (s *Service) processBillingEventV1(ctx context.Context, event *billing_event.BillingEventV1) error {
	eventID, _ := uuid.Parse(event.EventId) // no error can be here -- already validated

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("can not create tx: %w", err)
	}

	err = func() error {
		// check that it is a new message
		exists, err := tx.BillingOperations.Query().
			Where(billingoperations.UUID(eventID)).
			Exist(ctx)
		if err != nil {
			return fmt.Errorf("can not check event id: %w", err)
		}
		if exists {
			s.log.Debug("got message with same event id", zap.String("key", eventID.String()))
			return nil
		}

		userID, _ := uuid.Parse(event.UserId)
		u, err := tx.User.Query().
			Where(user.UUID(userID)).
			Only(ctx)
		if err != nil {
			return fmt.Errorf("can not find user for event")
		}

		var (
			operation billingoperations.Type
		)

		switch event.Event {
		case billing_event.Event_EVENT_PAYED_TO_USER:
			operation = billingoperations.TypePayout
		case billing_event.Event_EVENT_USER_BALANCE_BELOW_ZERO_AT_THE_END_OF_DAY:
			operation = billingoperations.TypeDebtMove
		}

		err = tx.BillingOperations.Create().
			SetUUID(eventID).
			SetAmount(event.Amount).
			SetType(operation).
			SetUserID(u.ID).
			SetTimestamp(event.Timestamp).
			Exec(ctx)

		return err
	}()

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("can not rollback tx: %w, err %w", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}
