package server

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lrmnt/AA6_homework/lib/auth"
	"net/http"
)

type claims struct {
	jwt.RegisteredClaims
	auth.UserInfo
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

	s.s.RespondJSON(w, cl.UserInfo)
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

	u, err := s.service.GetUserByName(r.Context(), name)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		RegisteredClaims: jwt.RegisteredClaims{},
		UserInfo: auth.UserInfo{
			ID:       u.ID,
			PublicID: u.UUID,
			Name:     u.Name,
			Role:     u.Edges.Role.Name,
		},
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		s.s.Respond401(w, "can not sign token", err)
		return
	}

	_, _ = w.Write([]byte(tokenString))
}
