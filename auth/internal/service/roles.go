package service

import (
	"context"
	"github.com/lrmnt/AA6_homework/auth/ent"
)

func (s *Service) ListRoles(ctx context.Context) ([]*ent.Role, error) {
	return s.client.Role.
		Query().
		All(ctx)
}

func (s *Service) CreateRole(ctx context.Context, name string) (*ent.Role, error) {
	return s.client.Role.
		Create().
		SetName(name).
		Save(ctx)
}
