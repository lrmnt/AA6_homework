package server

import (
	"github.com/lrmnt/AA6_homework/lib/auth"
	"net/http"
)

func (s *Server) listUserOwnBillingLog(w http.ResponseWriter, r *http.Request) {
	info, err := auth.GetUserInfo(r.Context())
	if err != nil {
		s.s.Respond500(w, "no user info", err)
		return
	}

	out, err := s.service.ListUserBillingLog(r.Context(), info.PublicID)
	if err != nil {
		s.s.Respond500(w, "can not load user log", err)
		return
	}

	s.s.RespondJSON(w, out)
}

func (s *Server) getOwnBalance(w http.ResponseWriter, r *http.Request) {
	info, err := auth.GetUserInfo(r.Context())
	if err != nil {
		s.s.Respond500(w, "no user info", err)
		return
	}

	out, err := s.service.GetUserBalance(r.Context(), info.PublicID)
	if err != nil {
		s.s.Respond500(w, "can not load user balance", err)
		return
	}

	s.s.RespondJSON(w, out)
}

func (s *Server) listUserOwnOperationsLog(w http.ResponseWriter, r *http.Request) {
	info, err := auth.GetUserInfo(r.Context())
	if err != nil {
		s.s.Respond500(w, "no user info", err)
		return
	}

	out, err := s.service.ListUserOperationsLog(r.Context(), info.PublicID)
	if err != nil {
		s.s.Respond500(w, "can not load user log", err)
		return
	}

	s.s.RespondJSON(w, out)
}

func (s *Server) getAdmimStats(w http.ResponseWriter, r *http.Request) {
	out, err := s.service.GetAdminStatistics(r.Context())
	if err != nil {
		s.s.Respond500(w, "can not load admin stats", err)
		return
	}

	s.s.RespondJSON(w, out)
}
