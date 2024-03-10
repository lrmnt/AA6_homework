package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lrmnt/AA6_homework/billing/internal/service/reports"
	"github.com/lrmnt/AA6_homework/lib/auth"
	"github.com/lrmnt/AA6_homework/lib/http"
	"go.uber.org/zap"
)

type Server struct {
	s       *http.Server
	log     *zap.Logger
	service *reports.Service
}

func New(authAddr, addr string, log *zap.Logger, service *reports.Service) *Server {
	router := chi.NewMux()

	s := &Server{
		s:       http.NewServer(addr, router, log),
		log:     log,
		service: service,
	}

	authClient := auth.New(authAddr)

	router.Use(middleware.Recoverer)

	router.With(authClient.AuthMiddleware()).Group(func(r chi.Router) {
		r.Get("/billing", s.listUserOwnBillingLog)
		r.Get("/operations", s.listUserOwnOperationsLog)
		r.Get("/balance", s.getOwnBalance)

		r.With(authClient.VerifyMiddleware("accountant", "admin")).
			Get("/stats", s.getAdmimStats)
	})

	return s
}

func (s *Server) Run() error {
	return s.s.Server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.s.Server.Shutdown(ctx)
}
