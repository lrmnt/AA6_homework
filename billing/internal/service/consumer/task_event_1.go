package consumer

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/billing/ent/operations"
	"github.com/lrmnt/AA6_homework/billing/ent/task"
	"github.com/lrmnt/AA6_homework/billing/ent/user"
	"github.com/lrmnt/AA6_homework/billing/internal/service/general"
	"github.com/lrmnt/AA6_homework/lib/api/schema"
	taskApi "github.com/lrmnt/AA6_homework/lib/api/schema/task_event"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func (s *Service) RunConsumeTaskEventV1(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			mes, err := s.taskEventV1Consumer.FetchMessage(ctx)
			if err != nil {
				s.log.Error("can not read message from queue", zap.Error(err))
				continue
			}

			var taskEvent taskApi.TaskEventV1

			err = proto.Unmarshal(mes.Value, &taskEvent)
			if err != nil {
				s.log.Error("can not unmarshal task event from queue", zap.Error(err))
				continue
			}

			ok, err := schema.ValidateTaskEventV1(&taskEvent)
			if err != nil || !ok { // invalid message format
				s.log.Warn("can not validate task event v1 message")

				err = s.taskEventV1Consumer.CommitMessages(ctx, mes)
				if err != nil {
					s.log.Error("can not commit task event", zap.Error(err))
				}

				continue
			}

			for i := 0; ; i++ {
				err = s.processTaskEventV1(ctx, &taskEvent)
				if err != nil {
					s.log.Warn("can not process task event", zap.Error(err), zap.Int("attempt", i+1))
					time.Sleep(retryTimeout * time.Duration(1<<i))
					continue
				}

				break
			}

			err = s.taskEventV1Consumer.CommitMessages(ctx, mes)
			if err != nil {
				s.log.Error("can not commit task event", zap.Error(err))
			}

			s.log.Debug("processed task event")
		}
	}
}

func (s *Service) processTaskEventV1(ctx context.Context, taskMessage *taskApi.TaskEventV1) error {
	eventID, _ := uuid.Parse(taskMessage.EventId) // no error can be here -- already validated

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("can not create tx: %w", err)
	}

	err = func() error {
		// check that it is a new message
		exists, err := tx.Operations.Query().
			Where(operations.UUID(eventID)).
			Exist(ctx)
		if err != nil {
			return fmt.Errorf("can not check event id: %w", err)
		}
		if exists {
			s.log.Debug("got message with same event id", zap.String("key", eventID.String()))
			return nil
		}

		userID, _ := uuid.Parse(taskMessage.AssigneeUserId)
		u, err := tx.User.Query().
			Where(user.UUID(userID)).
			Only(ctx)
		if err != nil {
			return fmt.Errorf("can not find user for task")
		}

		taskID, _ := uuid.Parse(taskMessage.TaskId)
		t, err := tx.Task.Query().
			Where(task.UUID(taskID)).
			Only(ctx)
		if err != nil {
			return fmt.Errorf("can not find task for event")
		}

		bc, err := general.GetCurrentBillingCycle(ctx, tx)
		if err != nil {
			return fmt.Errorf("can not get billing cycle")
		}

		var (
			operation operations.Type
			amount    int64
		)

		switch taskMessage.Event {
		case taskApi.Event_EVENT_DONE:
			operation = operations.TypeIncome
			amount = t.Cost
		case taskApi.Event_EVENT_REASSIGNED:
			operation = operations.TypeOutcome
			amount = rand.Int63() % 20
		}

		err = tx.Operations.Create().
			SetUUID(eventID).
			SetAmount(amount).
			SetType(operation).
			SetUserID(u.ID).
			SetTaskID(t.ID).
			SetBillingCycleID(bc.ID).
			SetTimestamp(time.Now()).
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
