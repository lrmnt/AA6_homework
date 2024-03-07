package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lrmnt/AA6_homework/lib/auth"
	"github.com/lrmnt/AA6_homework/lib/http"
	"github.com/lrmnt/AA6_homework/tasks/internal/service/tasks"
	"go.uber.org/zap"
)

type Server struct {
	s       *http.Server
	log     *zap.Logger
	service *tasks.Service
}

func New(authAddr, addr string, log *zap.Logger, service *tasks.Service) *Server {
	router := chi.NewMux()

	s := &Server{
		s:       http.NewServer(addr, router),
		log:     log,
		service: service,
	}

	authClient := auth.New(authAddr)

	router.Use(middleware.Recoverer)

	router.With(authClient.AuthMiddleware()).Group(func(r chi.Router) {
		r.Get("/tasks", s.listTasksForUser)
		r.Post("/tasks", s.createTask)
		r.Post("/tasks/{id}", s.updateTaskStatus)

		r.With(authClient.VerifyMiddleware("manager", "admin")).
			Post("/reassign", s.reassignTasks)
	})

	return s
}

func (s *Server) Run() error {
	return s.s.Server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.s.Server.Shutdown(ctx)
}
