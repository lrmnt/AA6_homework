package auth

import (
	"errors"
	"github.com/google/uuid"
)

const userCtxKey = "user"

var ErrNoUserInfo = errors.New("no user info")

type UserInfo struct {
	ID       int       `json:"id"`
	PublicID uuid.UUID `json:"public_id"`
	Name     string    `json:"name"`
	Role     string    `json:"role"`
}
