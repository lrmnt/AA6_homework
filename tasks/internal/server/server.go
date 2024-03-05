package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lrmnt/AA6_homework/tasks/ent"
	"github.com/lrmnt/AA6_homework/tasks/internal/kafka/producer"
	"go.uber.org/zap"
	"net/http"
)

const userCtxKey = "user"

type Server struct {
	s             *http.Server
	client        *ent.Client
	log           *zap.Logger
	tasksProducer *producer.Producer
}

func New(client *ent.Client, authAddr, addr string, l *zap.Logger, tasksProducer *producer.Producer) *Server {
	router := chi.NewMux()

	s := &Server{
		s: &http.Server{
			Handler: router,
			Addr:    addr,
		},
		client:        client,
		log:           l,
		tasksProducer: tasksProducer,
	}

	router.Use(middleware.Recoverer)

	router.With(AuthMiddleware(authAddr)).Group(func(r chi.Router) {
		r.Get("/tasks", s.listTasksForUser)
		r.Post("/tasks", s.createTask)
		r.Post("/tasks/{id}", s.updateTaskStatus)

		r.With(VerifyMiddleware("manager", "admin")).
			Post("/reassign", s.reassignTasks)
	})

	return s
}

func (s *Server) Run() error {
	return s.s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.s.Shutdown(ctx)
}

func (s *Server) responseError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	if err != nil {
		s.log.Error("error on http handle", zap.Error(err), zap.Int("code", code))
		_, _ = w.Write([]byte(err.Error()))
	}
}

func (s *Server) respondJSON(w http.ResponseWriter, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		s.log.Error("can not marshal response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(data)
}

func (s *Server) tx(ctx context.Context, fn func(tx *ent.Tx) error) error {
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
