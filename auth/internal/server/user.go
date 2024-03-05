package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	api "github.com/lrmnt/AA6_homework/auth/api/proto"
	"github.com/lrmnt/AA6_homework/auth/ent"
	"github.com/lrmnt/AA6_homework/auth/ent/user"
	"net/http"
	"strconv"
)

func (s *Server) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.client.User.Query().
		WithRole().
		All(r.Context())
	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	s.respondJSON(w, users)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	roleID, err := strconv.ParseInt(r.Form.Get("role"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var createdUser *ent.User
	err = s.tx(r.Context(), func(tx *ent.Tx) error {
		err := tx.User.Create().
			SetName(name).
			SetRoleID(int(roleID)).
			Exec(r.Context())
		if err != nil {
			return fmt.Errorf("can not create user in DB: %w", err)
		}

		createdUser, err = tx.User.Query().
			Where(user.Name(name)).
			WithRole().
			Only(r.Context())
		if err != nil {
			return fmt.Errorf("can not query user from DB: %w", err)
		}

		mes := &api.User{
			Action:         api.Action_ACTION_CREATED,
			PublicId:       createdUser.UUID.String(),
			Name:           createdUser.Name,
			Role:           createdUser.Edges.Role.Name,
			IdempotencyKey: uuid.New().String(),
		}

		data, err := proto.Marshal(mes)
		if err != nil {
			return fmt.Errorf("can not marshal message: %w", err)
		}

		err = s.userProducer.Produce(data)
		if err != nil {
			return fmt.Errorf("can not produce message: %w", err)
		}

		return nil
	})

	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	s.respondJSON(w, createdUser)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	roleID, err := strconv.ParseInt(r.Form.Get("role_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.tx(r.Context(), func(tx *ent.Tx) error {
		err := tx.User.Update().
			Where(user.ID(int(userID))).
			SetName(name).
			SetRoleID(int(roleID)).
			Exec(r.Context())
		if err != nil {
			return err
		}

		createdUser, err := tx.User.Query().
			Where(user.Name(name)).
			Only(r.Context())
		if err != nil {
			return err
		}

		mes := &api.User{
			Action:         api.Action_ACTION_MODIFIED,
			PublicId:       createdUser.UUID.String(),
			Name:           createdUser.Name,
			Role:           createdUser.Edges.Role.Name,
			IdempotencyKey: uuid.New().String(),
		}

		data, err := proto.Marshal(mes)
		if err != nil {
			return err
		}

		err = s.userProducer.Produce(data)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		s.responseError(w, http.StatusInternalServerError, err)
		return
	}

	_, _ = w.Write([]byte("ok"))
}
