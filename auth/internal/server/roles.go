package server

import (
	"net/http"
)

func (s *Server) listRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := s.service.ListRoles(r.Context())
	if err != nil {
		s.s.Respond500(w, "can not list roles", err)
		return
	}

	s.s.RespondJSON(w, roles)
}

func (s *Server) createRole(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.s.Respond400(w, "can not parse form", err)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		s.s.Respond400(w, "no name", nil)
	}

	role, err := s.service.CreateRole(r.Context(), name)
	if err != nil {
		s.s.Respond500(w, "can not create role", err)
		return
	}

	s.s.RespondJSON(w, role)
}
