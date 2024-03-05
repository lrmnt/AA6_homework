package model

import "github.com/google/uuid"

type UserInfo struct {
	ID       int       `json:"id"`
	PublicID uuid.UUID `json:"public_id"`
	Name     string    `json:"name"`
	Role     string    `json:"role"`
}
