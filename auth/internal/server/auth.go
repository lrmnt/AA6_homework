package server

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lrmnt/AA6_homework/auth/ent/user"
	"github.com/lrmnt/AA6_homework/auth/pkg/model"
	"net/http"
)

type claims struct {
	jwt.RegisteredClaims
	model.UserInfo
}

func (s *Server) validate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t := r.Form.Get("token")

	token, err := jwt.ParseWithClaims(t, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cl, ok := token.Claims.(*claims)
	if !ok || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	s.respondJSON(w, cl.UserInfo)
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := r.Form.Get("user")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := s.client.User.Query().
		Where(user.Name(name)).
		WithRole().
		Only(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		RegisteredClaims: jwt.RegisteredClaims{},
		UserInfo: model.UserInfo{
			ID:       u.ID,
			PublicID: u.UUID,
			Name:     u.Name,
			Role:     u.Edges.Role.Name,
		},
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		s.responseError(w, http.StatusUnauthorized, fmt.Errorf("can not sign token: %w", err))
		return
	}

	_, _ = w.Write([]byte(tokenString))
}
