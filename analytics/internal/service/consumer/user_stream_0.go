package consumer

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/analytics/ent/user"
	"github.com/lrmnt/AA6_homework/analytics/ent/userlog"
	userApi "github.com/lrmnt/AA6_homework/lib/api/schema/user"
	"go.uber.org/zap"
)

func (s *Service) RunConsumeUserMessageV0(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			mes, err := s.userStreamV0Consumer.FetchMessage(ctx)
			if err != nil {
				s.log.Error("can not read message from queue", zap.Error(err))
				continue
			}

			var userMessage userApi.User

			err = proto.Unmarshal(mes.Value, &userMessage)
			if err != nil {
				s.log.Error("can not unmarshal user message from queue", zap.Error(err))
				continue
			}

			err = s.processUserMessageV0(ctx, &userMessage)
			if err != nil {
				s.log.Error("can not process user message", zap.Error(err))
				continue
			}

			err = s.userStreamV0Consumer.CommitMessages(ctx, mes)
			if err != nil {
				s.log.Error("can not commit user message", zap.Error(err))
			}

		}
	}
}

func (s *Service) processUserMessageV0(ctx context.Context, userMessage *userApi.User) error {
	publicID, err := uuid.Parse(userMessage.PublicId)
	if err != nil {
		return fmt.Errorf("can not parse public id: %w", err)
	}

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("can not create tx: %w", err)
	}

	err = func() error {
		// check that it is a new message
		exists, err := tx.UserLog.Query().
			Where(userlog.IdempotencyKey(userMessage.IdempotencyKey)).
			Exist(ctx)
		if err != nil {
			return fmt.Errorf("can not check idempotency key: %w", err)
		}
		if exists {
			s.log.Debug("got message with same idempotency key", zap.String("key", userMessage.IdempotencyKey))
			return nil
		}

		// add idempotency key to db
		err = tx.UserLog.Create().SetIdempotencyKey(userMessage.IdempotencyKey).Exec(ctx)
		if err != nil {
			return fmt.Errorf("can not store idempotency key in db: %w", err)
		}

		switch userMessage.Action {
		case userApi.Action_ACTION_CREATED:
			err = tx.User.Create().
				SetName(userMessage.Name).
				SetRole(userMessage.Role).
				SetUUID(publicID).
				Exec(ctx)

			if err != nil {
				return fmt.Errorf("can not save user: %w", err)
			}

		case userApi.Action_ACTION_MODIFIED:
			err = s.client.User.Update().
				Where(user.UUID(publicID)).
				SetName(userMessage.Name).
				SetRole(userMessage.Role).
				Exec(ctx)

			if err != nil {
				return fmt.Errorf("can not update user: %w", err)
			}

		default:
			return errors.New("unknown action")
		}

		return nil
	}()

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("can not rollback tx: %w, err %w", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}
