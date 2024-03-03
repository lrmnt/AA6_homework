package server

import (
	"context"
	"encoding/json"
	"github.com/lrmnt/AA6_homework/auth/pkg/model"
	"io"
	"net/http"
	"strings"
)

func AuthMiddleware(addr string) func(http.Handler) http.Handler {
	client := http.Client{}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parts := strings.Split(r.Header.Get("Authorization"), " ")
			if len(parts) != 2 || parts[0] != "BEARER" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, addr+"/validate?token="+parts[1], nil)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			defer resp.Body.Close()

			var userInfo model.UserInfo

			err = json.Unmarshal(data, &userInfo)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), userCtxKey, &userInfo))

			next.ServeHTTP(w, r)
		})
	}
}

func VerifyMiddleware(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(userCtxKey).(*model.UserInfo)
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			for _, role := range roles {
				if role == user.Role {
					next.ServeHTTP(w, r)
					return
				}
			}

			w.WriteHeader(http.StatusForbidden)
		})
	}
}
