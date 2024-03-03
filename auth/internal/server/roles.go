package server

import "net/http"

func (s *Server) listRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := s.client.Role.
		Query().
		All(r.Context())
	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	s.respondJSON(w, roles)
}

func (s *Server) createRole(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	role, err := s.client.Role.
		Create().
		SetName(name).
		Save(r.Context())
	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	s.respondJSON(w, role)
}
