package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (s *Server) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.service.ListUsers(r.Context())
	if err != nil {
		s.s.Respond500(w, "can nol list users", err)
		return
	}

	s.s.RespondJSON(w, users)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.s.Respond400(w, "can not parse form", err)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		s.s.Respond400(w, "no name", nil)
		return
	}

	roleID, err := strconv.ParseInt(r.Form.Get("role"), 10, 64)
	if err != nil {
		s.s.Respond400(w, "no valid role id", err)
		return
	}

	createdUser, err := s.service.CreateUser(r.Context(), name, int(roleID))
	if err != nil {
		s.s.Respond500(w, "can not create user", err)
		return
	}

	s.s.RespondJSON(w, createdUser)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.s.Respond400(w, "can not parse form", err)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		s.s.Respond400(w, "no name", nil)
		return
	}

	userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		s.s.Respond400(w, "no valid user id", err)
		return
	}

	roleID, err := strconv.ParseInt(r.Form.Get("role_id"), 10, 64)
	if err != nil {
		s.s.Respond400(w, "no valid role id", err)
		return
	}

	updatedUser, err := s.service.UpdateUser(r.Context(), name, int(roleID), int(userID))

	if err != nil {
		s.s.Respond500(w, "can not update user", err)
		return
	}

	s.s.RespondJSON(w, updatedUser)
}
