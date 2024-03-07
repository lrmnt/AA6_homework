package auth

import (
	"context"
	"net/http"
	"strings"
)

func (c *Client) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parts := strings.Split(r.Header.Get("Authorization"), " ")
			if len(parts) != 2 || parts[0] != "BEARER" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			userInfo, err := c.validate(r.Context(), parts[1])
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(err.Error()))
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), userCtxKey, &userInfo))

			next.ServeHTTP(w, r)
		})
	}
}

func (c *Client) VerifyMiddleware(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := GetUserInfo(r.Context())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(err.Error()))
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

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	userInfo, ok := ctx.Value(userCtxKey).(*UserInfo)
	if !ok {
		return nil, ErrNoUserInfo
	}

	return userInfo, nil
}
