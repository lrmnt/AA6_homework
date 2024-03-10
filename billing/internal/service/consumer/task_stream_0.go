package consumer

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/billing/ent/task"
	"github.com/lrmnt/AA6_homework/billing/ent/tasklog"
	"github.com/lrmnt/AA6_homework/billing/ent/user"
	taskApi "github.com/lrmnt/AA6_homework/lib/api/schema/task"
	"go.uber.org/zap"
)

func (s *Service) RunConsumeTaskMessageV0(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			mes, err := s.taskStreamV0Consumer.FetchMessage(ctx)
			if err != nil {
				s.log.Error("can not read message from queue", zap.Error(err))
				continue
			}

			var taskMessage taskApi.Task

			err = proto.Unmarshal(mes.Value, &taskMessage)
			if err != nil {
				s.log.Error("can not unmarshal task message from queue", zap.Error(err))
				continue
			}

			err = s.processTaskMessageV0(ctx, &taskMessage)
			if err != nil {
				s.log.Error("can not process task message", zap.Error(err))
				continue
			}

			err = s.taskStreamV0Consumer.CommitMessages(ctx, mes)
			if err != nil {
				s.log.Error("can not commit task message", zap.Error(err))
			}

		}
	}

}

func (s *Service) processTaskMessageV0(ctx context.Context, taskMessage *taskApi.Task) error {
	publicID, err := uuid.Parse(taskMessage.PublicId)
	if err != nil {
		return fmt.Errorf("can not parse public id: %w", err)
	}

	userID, err := uuid.Parse(taskMessage.UserId)
	if err != nil {
		return fmt.Errorf("can not parse user id: %w", err)
	}

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("can not create tx: %w", err)
	}

	err = func() error {
		// check that it is a new message
		exists, err := tx.TaskLog.Query().
			Where(tasklog.IdempotencyKey(taskMessage.IdempotencyKey)).
			Exist(ctx)
		if err != nil {
			return fmt.Errorf("can not check idempotency key: %w", err)
		}
		if exists {
			s.log.Debug("got message with same idempotency key", zap.String("key", taskMessage.IdempotencyKey))
			return nil
		}

		// add idempotency key to db
		err = tx.TaskLog.Create().SetIdempotencyKey(taskMessage.IdempotencyKey).Exec(ctx)
		if err != nil {
			return fmt.Errorf("can not store idempotency key in db: %w", err)
		}

		u, err := tx.User.Query().
			Where(user.UUID(userID)).
			Only(ctx)
		if err != nil {
			return err
		}

		switch taskMessage.Action {
		case taskApi.Action_ACTION_CREATED:
			err = tx.Task.Create().
				SetUUID(publicID).
				SetTitle(taskMessage.Title).
				SetDescription(taskMessage.Description).
				SetCost(taskMessage.Cost).
				SetUser(u).
				Exec(ctx)
			if err != nil {
				return fmt.Errorf("can not save task: %w", err)
			}

			if err != nil {
				return fmt.Errorf("can not save task v0: %w", err)
			}

		case taskApi.Action_ACTION_MODIFIED:
			err = s.client.Task.Update().
				Where(task.UUID(publicID)).
				SetTitle(taskMessage.Title).
				SetDescription(taskMessage.Description).
				SetCost(taskMessage.Cost).
				SetUser(u).
				Exec(ctx)

			if err != nil {
				return fmt.Errorf("can not update task v0: %w", err)
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
