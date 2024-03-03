package server

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/auth/pkg/model"
	api "github.com/lrmnt/AA6_homework/tasks/api/proto"
	"github.com/lrmnt/AA6_homework/tasks/ent"
	"github.com/lrmnt/AA6_homework/tasks/ent/task"
	"github.com/lrmnt/AA6_homework/tasks/ent/user"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strconv"
)

var (
	errNotOwnTask = errors.New("not own task")
)

func (s *Server) listTasksForUser(w http.ResponseWriter, r *http.Request) {
	info, ok := r.Context().Value(userCtxKey).(*model.UserInfo)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tasks, err := s.client.Task.
		Query().
		Where(task.HasUserWith(user.UUID(info.PublicID))).
		WithUser().
		All(r.Context())
	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	s.respondJSON(w, tasks)
}

func statusFromString(s string) (api.Status, error) {
	switch s {
	case "done":
		return api.Status_STATUS_DONE, nil
	case "in_progress":
		return api.Status_STATUS_IN_PROGRESS, nil
	case "todo":
		return api.Status_STATUS_TODO, nil
	}

	return api.Status_STATUS_UNKNOWN, errors.New("unknown status")
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	title := r.Form.Get("title")
	description := r.Form.Get("description")

	var createdTask *ent.Task
	err = s.tx(r.Context(), func(tx *ent.Tx) error {
		id, err := getRandomUserId(r.Context(), tx)
		if err != nil {
			return err
		}

		createdTask, err = tx.Task.Create().
			SetUserID(id).
			SetStatus(task.DefaultStatus).
			SetDescription(description).
			SetTitle(title).
			SetPrice(20 + rand.Int63()%20).
			Save(r.Context())
		if err != nil {
			return err
		}

		status, err := statusFromString(createdTask.Status.String())
		if err != nil {
			return err
		}

		mes := &api.Task{
			Action:         api.Action_ACTION_CREATED,
			Status:         status,
			PublicId:       createdTask.UUID.String(),
			Title:          createdTask.Title,
			Description:    createdTask.Description,
			Cost:           createdTask.Price,
			UserId:         int64(id),
			IdempotencyKey: uuid.New().String(),
		}

		data, err := proto.Marshal(mes)
		if err != nil {
			return err
		}

		if err = s.tasksProducer.Produce(data); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
	}

	s.respondJSON(w, createdTask)
}

func (s *Server) updateTaskStatus(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := r.Context().Value(userCtxKey).(*model.UserInfo)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	taskID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		s.responseError(w, http.StatusBadRequest, err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		s.responseError(w, http.StatusBadRequest, err)
		return
	}

	status := r.Form.Get("status")

	var updatedTask *ent.Task
	err = s.tx(r.Context(), func(tx *ent.Tx) error {
		taskToUpdate, err := tx.Task.Query().
			Where(task.ID(int(taskID))).
			WithUser().
			Only(r.Context())
		if err != nil {
			return err
		}

		if taskToUpdate.Edges.User.UUID != userInfo.PublicID { // we can change statuses only for own tasks
			return errNotOwnTask
		}

		updatedTask, err = tx.Task.UpdateOneID(int(taskID)).
			SetStatus(task.Status(status)). // todo check status name
			Save(r.Context())
		if err != nil {
			return err
		}

		apiStatus, err := statusFromString(updatedTask.Status.String())
		if err != nil {
			return err
		}

		mes := &api.Task{
			Action:         api.Action_ACTION_MODIFIED,
			Status:         apiStatus,
			PublicId:       updatedTask.UUID.String(),
			Title:          updatedTask.Title,
			Description:    updatedTask.Description,
			Cost:           updatedTask.Price,
			UserId:         int64(userInfo.ID),
			IdempotencyKey: uuid.New().String(),
		}

		data, err := proto.Marshal(mes)
		if err != nil {
			return err
		}

		if err = s.tasksProducer.Produce(data); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, errNotOwnTask) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	s.respondJSON(w, updatedTask)
}

func (s *Server) reassignTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.client.Task.Query().
		Where(task.StatusNotIn("done")).
		All(r.Context())
	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	for _, t := range tasks {
		err = s.reassignTaskRandomly(r.Context(), t.ID)
		if err != nil {
			s.log.Error("can not reassign task", zap.Error(err), zap.Int("id", t.ID))
		}
	}

	_, _ = w.Write([]byte("ok"))
}

func (s *Server) reassignTaskRandomly(ctx context.Context, taskID int) error {
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

		status, err := statusFromString(loadedTask.Status.String())
		if err != nil {
			return err
		}

		// selecting random user inside of transaction to be sure that user exists and have correct role
		newID, err := getRandomUserId(ctx, tx)
		if err != nil {
			return err
		}

		if newID != loadedTask.Edges.User.ID { // new user to assign
			err = tx.Task.UpdateOneID(loadedTask.ID).
				SetUserID(newID).
				Exec(ctx)
			if err != nil {
				return err
			}

			mes := &api.Task{
				Action:         api.Action_ACTION_REASSIGNED,
				Status:         status,
				PublicId:       loadedTask.UUID.String(),
				Title:          loadedTask.Title,
				Description:    loadedTask.Description,
				Cost:           loadedTask.Price,
				UserId:         int64(newID),
				IdempotencyKey: uuid.New().String(),
			}

			data, err := proto.Marshal(mes)
			if err != nil {
				return err
			}

			err = s.tasksProducer.Produce(data)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func getRandomUserId(ctx context.Context, tx *ent.Tx) (int, error) {
	users, err := tx.User.Query().
		Where(user.RoleNotIn("manager", "admin")).
		All(ctx)

	if err != nil {
		return 0, err
	}

	ids := make([]int, len(users))
	for i := range users {
		ids[i] = users[i].ID
	}

	return ids[rand.Int()%len(ids)], nil
}
