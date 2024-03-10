package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lrmnt/AA6_homework/auth/internal/service"
	"github.com/lrmnt/AA6_homework/lib/http"
	"go.uber.org/zap"
)

type Server struct {
	s         *http.Server
	log       *zap.Logger
	jwtSecret string
	service   *service.Service
}

func New(
	jwtSecret string,
	addr string,
	log *zap.Logger,
	service *service.Service,
) (*Server, error) {
	router := chi.NewMux()

	s := &Server{
		s:         http.NewServer(addr, router, log),
		jwtSecret: jwtSecret,
		service:   service,
		log:       log,
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
	return s.s.Server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.s.Server.Shutdown(ctx)
}
