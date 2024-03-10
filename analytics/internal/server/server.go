package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lrmnt/AA6_homework/analytics/internal/service/reports"
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

	router.With(authClient.AuthMiddleware()).
		With(authClient.VerifyMiddleware("admin")).
		Group(func(r chi.Router) {
			r.Get("/last_day_sum", s.getLastDaySum)
			r.Get("/last_day_debts", s.getLastDayPopugDebtsCount)
			r.Get("/highest_price", s.getHighestPriceForPeriod)
		})

	return s
}

func (s *Server) Run() error {
	return s.s.Server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.s.Server.Shutdown(ctx)
}
