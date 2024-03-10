package http

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	Server *http.Server
	log    *zap.Logger
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewServer(addr string, router *chi.Mux, log *zap.Logger) *Server {
	return &Server{
		Server: &http.Server{
			Handler: router,
			Addr:    addr,
		},
		log: log,
	}
}

func (s *Server) Respond500(w http.ResponseWriter, mes string, err error) {
	s.responseError(w, http.StatusInternalServerError, mes, err)
}

func (s *Server) Respond400(w http.ResponseWriter, mes string, err error) {
	s.responseError(w, http.StatusBadRequest, mes, err)
}

func (s *Server) Respond401(w http.ResponseWriter, mes string, err error) {
	s.responseError(w, http.StatusUnauthorized, mes, err)
}

func (s *Server) Respond403(w http.ResponseWriter, mes string, err error) {
	s.responseError(w, http.StatusForbidden, mes, err)
}

func (s *Server) responseError(w http.ResponseWriter, code int, mes string, err error) {
	w.WriteHeader(code)
	if code == http.StatusInternalServerError || err != nil {
		s.log.Error(mes, zap.Error(err), zap.Int("code", code))
	}

	out := errorResponse{
		Message: mes,
	}
	if err != nil {
		out.Error = err.Error()
	}

	s.RespondJSON(w, out)
}

func (s *Server) RespondJSON(w http.ResponseWriter, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		s.log.Error("can not marshal response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(data)
}
