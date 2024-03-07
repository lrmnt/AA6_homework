package server

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/lrmnt/AA6_homework/lib/auth"
	"github.com/lrmnt/AA6_homework/tasks/ent/task"
	"github.com/lrmnt/AA6_homework/tasks/internal/service/tasks"
	"net/http"
	"strconv"
)

func (s *Server) listTasksForUser(w http.ResponseWriter, r *http.Request) {
	info, err := auth.GetUserInfo(r.Context())
	if err != nil {
		s.s.Respond500(w, "no user info", err)
		return
	}

	tasks, err := s.service.ListTasksForUser(r.Context(), info.PublicID)
	if err != nil {
		s.s.Respond500(w, "can not list tasks", err)
		return
	}

	s.s.RespondJSON(w, tasks)
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.s.Respond400(w, "can not parse form", err)
		return
	}

	title := r.Form.Get("title")
	description := r.Form.Get("description")

	if title == "" || description == "" {
		s.s.Respond400(w, "no titles or description", nil)
		return
	}

	createdTask, err := s.service.CreateTask(r.Context(), title, description)

	if err != nil {
		s.s.Respond500(w, "can not create task", err)
	}

	s.s.RespondJSON(w, createdTask)
}

func (s *Server) updateTaskStatus(w http.ResponseWriter, r *http.Request) {
	userInfo, err := auth.GetUserInfo(r.Context())
	if err != nil {
		s.s.Respond500(w, "no user info", err)
		return
	}

	taskID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		s.s.Respond400(w, "no valid task id", err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		s.s.Respond400(w, "can not parse form", err)
		return
	}

	status := r.Form.Get("status")
	switch status {
	case task.StatusDone.String(), task.StatusTodo.String(), task.StatusInProgress.String():
	default:
		s.s.Respond400(w, "no valid status name", err)
		return
	}

	updatedTask, err := s.service.UpdateTaskStatus(r.Context(), int(taskID), userInfo.PublicID, task.Status(status))
	if err != nil {
		if errors.Is(err, tasks.ErrNotOwnTask) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		s.s.Respond500(w, "can not update task", err)
		return
	}
	s.s.RespondJSON(w, updatedTask)
}

func (s *Server) reassignTasks(w http.ResponseWriter, r *http.Request) {
	err := s.service.ReassignTasks(r.Context())
	if err != nil {
		s.s.Respond500(w, "can not reassign tasks", err)
		return
	}

	_, _ = w.Write([]byte("ok"))
}
