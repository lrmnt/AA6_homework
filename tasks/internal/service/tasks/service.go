package tasks

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/lib/api/schema"
	"github.com/lrmnt/AA6_homework/lib/api/schema/task_event"
	"github.com/lrmnt/AA6_homework/lib/api/schema/task_stream"
	"github.com/lrmnt/AA6_homework/lib/kafka"
	"github.com/lrmnt/AA6_homework/tasks/ent"
	"github.com/lrmnt/AA6_homework/tasks/ent/task"
	"github.com/lrmnt/AA6_homework/tasks/ent/user"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

var (
	ErrNotOwnTask = errors.New("not own task")
)

type Service struct {
	log                 *zap.Logger
	client              *ent.Client
	tasksStreamProducer *kafka.SaramaProducer
	tasksEventProducer  *kafka.SaramaProducer
}

func New(
	log *zap.Logger,
	client *ent.Client,
	tasksStreamProducer *kafka.SaramaProducer,
	tasksEvenProducer *kafka.SaramaProducer,
) *Service {
	return &Service{
		log:                 log,
		client:              client,
		tasksStreamProducer: tasksStreamProducer,
		tasksEventProducer:  tasksEvenProducer,
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

func (s *Service) ListTasksForUser(ctx context.Context, id uuid.UUID) ([]*ent.Task, error) {
	return s.client.Task.
		Query().
		Where(task.HasUserWith(user.UUID(id))).
		WithUser().
		All(ctx)
}

func statusFromString(s string) (task_stream.Status, error) {
	switch s {
	case "done":
		return task_stream.Status_STATUS_DONE, nil
	case "in_progress":
		return task_stream.Status_STATUS_IN_PROGRESS, nil
	case "todo":
		return task_stream.Status_STATUS_TODO, nil
	}

	return task_stream.Status_STATUS_UNKNOWN, errors.New("unknown status")
}

func getRandomUser(ctx context.Context, tx *ent.Tx) (*ent.User, error) {
	users, err := tx.User.Query().
		Where(user.RoleNotIn("manager", "admin")).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return users[rand.Int()%len(users)], nil
}

func (s *Service) CreateTask(ctx context.Context, title, description string) (*ent.Task, error) {
	var createdTask *ent.Task
	err := s.tx(ctx, func(tx *ent.Tx) error {
		randomUser, err := getRandomUser(ctx, tx)
		if err != nil {
			return fmt.Errorf("can not get user to assign: %w", err)
		}

		createdTask, err = tx.Task.Create().
			SetUserID(randomUser.ID).
			SetStatus(task.DefaultStatus).
			SetDescription(description).
			SetTitle(title).
			SetPrice(20 + rand.Int63()%20).
			Save(ctx)
		if err != nil {
			return err
		}

		status, err := statusFromString(createdTask.Status.String())
		if err != nil {
			return err
		}

		mes := &task_stream.TaskStreamV1{
			Action:         task_stream.Action_ACTION_CREATED,
			Status:         status,
			PublicId:       createdTask.UUID.String(),
			Title:          createdTask.Title,
			Description:    createdTask.Description,
			Cost:           createdTask.Price,
			UserId:         randomUser.UUID.String(),
			IdempotencyKey: uuid.New().String(),
			Timestamp:      time.Now().UnixNano(),
		}

		_, err = schema.ValidateTaskStreamV1(mes)
		if err != nil {
			return fmt.Errorf("can not validate message: %w", err)
		}

		data, err := proto.Marshal(mes)
		if err != nil {
			return fmt.Errorf("can not marshal message: %w", err)
		}

		if err = s.tasksStreamProducer.Send(data); err != nil {
			return fmt.Errorf("can not produce stream message: %w", err)
		}

		if err = s.produceTaskEvent(task_event.Event_EVENT_REASSIGNED, randomUser.UUID, createdTask.UUID); err != nil {
			return fmt.Errorf("can not produce event message: %w", err)
		}

		return nil
	})

	return createdTask, err
}

func (s *Service) UpdateTaskStatus(ctx context.Context, taskID int, userID uuid.UUID, status task.Status) (*ent.Task, error) {
	var updatedTask *ent.Task
	err := s.tx(ctx, func(tx *ent.Tx) error {
		taskToUpdate, err := tx.Task.Query().
			Where(task.ID(taskID)).
			WithUser().
			Only(ctx)
		if err != nil {
			return err
		}

		if taskToUpdate.Edges.User.UUID != userID { // we can change statuses only for own tasks
			return ErrNotOwnTask
		}

		updatedTask, err = tx.Task.UpdateOneID(taskID).
			SetStatus(status).
			Save(ctx)
		if err != nil {
			return err
		}

		apiStatus, err := statusFromString(updatedTask.Status.String())
		if err != nil {
			return err
		}

		if apiStatus == task_stream.Status_STATUS_DONE {
			if err = s.produceTaskEvent(task_event.Event_EVENT_DONE, userID, updatedTask.UUID); err != nil {
				return err
			}
		}

		return nil
	})

	return updatedTask, err
}

func (s *Service) ReassignTasks(ctx context.Context) error {
	tasks, err := s.client.Task.Query().
		Where(task.StatusNotIn(task.StatusDone)).
		All(ctx)
	if err != nil {
		return err
	}

	for _, t := range tasks {
		err = s.reassignTaskRandomly(ctx, t.ID)
		if err != nil {
			s.log.Error("can not reassign task", zap.Error(err), zap.Int("id", t.ID))
		}
	}

	return nil
}

func (s *Service) reassignTaskRandomly(ctx context.Context, taskID int) error {
	return s.tx(ctx, func(tx *ent.Tx) error {

		// we loaded tasks list outside of transaction, so some of them may change status between load and reassign.
		// so load task by id again to check status
		loadedTask, err := tx.Task.Query().
			Where(task.ID(taskID)).
			WithUser().
			Only(ctx)
		if err != nil {
			return err
		}

		if loadedTask.Status == "done" {
			s.log.Debug("task changed status before reassign", zap.Int("id", taskID))
			return nil
		}

		// selecting random user inside of transaction to be sure that user exists and have correct role
		randomUser, err := getRandomUser(ctx, tx)
		if err != nil {
			return err
		}

		if randomUser.ID != loadedTask.Edges.User.ID { // new user to assign
			err = tx.Task.UpdateOneID(loadedTask.ID).
				SetUserID(randomUser.ID).
				Exec(ctx)
			if err != nil {
				return err
			}

			if err = s.produceTaskEvent(task_event.Event_EVENT_REASSIGNED, randomUser.UUID, loadedTask.UUID); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *Service) produceTaskEvent(event task_event.Event, userID, taskID uuid.UUID) error {
	mes := &task_event.TaskEventV1{
		Event:          event,
		Timestamp:      time.Now().UnixNano(),
		EventId:        uuid.New().String(),
		AssigneeUserId: userID.String(),
		TaskId:         taskID.String(),
	}

	_, err := schema.ValidateTaskEventV1(mes)
	if err != nil {
		return err
	}

	data, err := proto.Marshal(mes)
	if err != nil {
		return err
	}

	return s.tasksEventProducer.Send(data)
}
