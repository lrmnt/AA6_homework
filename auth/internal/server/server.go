package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lrmnt/AA6_homework/auth/ent"
	"github.com/lrmnt/AA6_homework/auth/internal/producer"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	s            *http.Server
	l            *zap.Logger
	jwtSecret    string
	client       *ent.Client
	userProducer *producer.Producer
}

func New(
	client *ent.Client,
	jwtSecret string,
	addr string,
	l *zap.Logger,
	userProducer *producer.Producer,
) (*Server, error) {
	router := chi.NewMux()

	s := &Server{
		s: &http.Server{
			Handler: router,
			Addr:    addr,
		},
		client:       client,
		jwtSecret:    jwtSecret,
		l:            l,
		userProducer: userProducer,
	}

	router.Use(middleware.Recoverer)

	// roles CRUD
	router.Get("/roles", s.listRoles)
	router.Post("/roles", s.createRole)

	// users CRUD
	router.Get("/users", s.listUsers)
	router.Post("/users", s.createUser)
	router.Post("/users/{id}", s.updateUser)

	// auth
	router.Post("/login", s.login)
	router.Post("/validate", s.validate)

	return s, nil
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
		s.l.Error("error on http handle", zap.Error(err), zap.Int("code", code))
		_, _ = w.Write([]byte(err.Error()))
	}
}

func (s *Server) respondJSON(w http.ResponseWriter, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		s.l.Error("can not marshal response", zap.Error(err))
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
			return fmt.Errorf("can not rollback tx: %w, Err: %w", rbErr, err)
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("can not commit tx: %w", err)
	}

	return nil
}
