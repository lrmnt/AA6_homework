package consumer

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/analytics/ent/task"
	"github.com/lrmnt/AA6_homework/analytics/ent/tasklog"
	"github.com/lrmnt/AA6_homework/analytics/ent/user"
	"github.com/lrmnt/AA6_homework/lib/api/schema"
	taskApi "github.com/lrmnt/AA6_homework/lib/api/schema/task_stream"
	"go.uber.org/zap"
)

func (s *Service) RunConsumeTaskMessageV1(ctx context.Context) error {

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			mes, err := s.taskStreamV1Consumer.FetchMessage(ctx)
			if err != nil {
				s.log.Error("can not read message from queue", zap.Error(err))
				continue
			}

			var taskMessage taskApi.TaskStreamV1

			err = proto.Unmarshal(mes.Value, &taskMessage)
			if err != nil {
				s.log.Error("can not unmarshal user message from queue", zap.Error(err))
				continue
			}

			ok, err := schema.ValidateTaskStreamV1(&taskMessage)
			if err != nil || !ok { // invalid message format
				s.log.Warn("can not validate user stream v1 message")

				err = s.userStreamV0Consumer.CommitMessages(ctx, mes)
				if err != nil {
					s.log.Error("can not commit user message", zap.Error(err))
				}

				continue
			}

			err = s.processTaskMessageV1(ctx, &taskMessage)
			if err != nil {
				s.log.Error("can not process user message", zap.Error(err))
				continue
			}

			err = s.taskStreamV1Consumer.CommitMessages(ctx, mes)
			if err != nil {
				s.log.Error("can not commit user message", zap.Error(err))
			}

		}
	}

}

func (s *Service) processTaskMessageV1(ctx context.Context, taskMessage *taskApi.TaskStreamV1) error {
	publicID, _ := uuid.Parse(taskMessage.PublicId) // no error can be here -- already validated

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

		userID, _ := uuid.Parse(taskMessage.UserId)

		u, err := tx.User.Query().
			Where(user.UUID(userID)).
			Only(ctx)
		if err != nil {
			return fmt.Errorf("can not find user for task")
		}

		switch taskMessage.Action {
		case taskApi.Action_ACTION_CREATED:
			err = tx.Task.Create().
				SetUUID(publicID).
				SetTitle(taskMessage.Title).
				SetDescription(taskMessage.Description).
				SetCost(taskMessage.Cost).
				SetUser(u).
				SetJiraID(taskMessage.JiraId).
				SetTimestamp(taskMessage.Timestamp).
				Exec(ctx)
			if err != nil {
				return fmt.Errorf("can not save task: %w", err)
			}

		case taskApi.Action_ACTION_MODIFIED:
			err = s.client.Task.Update().
				Where(task.UUID(publicID)).
				SetTitle(taskMessage.Title).
				SetDescription(taskMessage.Description).
				SetCost(taskMessage.Cost).
				SetUser(u).
				SetJiraID(taskMessage.JiraId).
				Exec(ctx)

			if err != nil {
				return fmt.Errorf("can not update task: %w", err)
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
