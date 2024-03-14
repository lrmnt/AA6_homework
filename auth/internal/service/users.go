package service

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/lrmnt/AA6_homework/auth/ent"
	"github.com/lrmnt/AA6_homework/auth/ent/user"
	"github.com/lrmnt/AA6_homework/lib/api/schema"
	"github.com/lrmnt/AA6_homework/lib/api/schema/user_stream"
	"time"
)

func (s *Service) ListUsers(ctx context.Context) ([]*ent.User, error) {
	return s.client.User.Query().
		WithRole().
		All(ctx)
}

func (s *Service) GetUserByName(ctx context.Context, name string) (*ent.User, error) {
	return s.client.User.Query().
		Where(user.Name(name)).
		WithRole().
		Only(ctx)
}

func (s *Service) CreateUser(ctx context.Context, name string, roleID int) (*ent.User, error) {
	var createdUser *ent.User

	err := s.tx(ctx, func(tx *ent.Tx) error {
		err := tx.User.Create().
			SetName(name).
			SetRoleID(int(roleID)).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("can not create user in DB: %w", err)
		}

		createdUser, err = tx.User.Query().
			Where(user.Name(name)).
			WithRole().
			Only(ctx)
		if err != nil {
			return fmt.Errorf("can not query user from DB: %w", err)
		}

		mes := &user_stream.UserStreamV1{
			Action:         user_stream.Action_ACTION_CREATED,
			PublicId:       createdUser.UUID.String(),
			Name:           createdUser.Name,
			Role:           createdUser.Edges.Role.Name,
			IdempotencyKey: uuid.New().String(),
			Timestamp:      time.Now().UnixNano(),
		}

		_, err = schema.ValidateUserStreamV1(mes)
		if err != nil {
			return err
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

	return createdUser, err
}

func (s *Service) UpdateUser(ctx context.Context, name string, roleID, userID int) (*ent.User, error) {
	var updatedUser *ent.User

	err := s.tx(ctx, func(tx *ent.Tx) error {
		err := tx.User.Update().
			Where(user.ID(userID)).
			SetName(name).
			SetRoleID(roleID).
			Exec(ctx)
		if err != nil {
			return err
		}

		updatedUser, err = tx.User.Query().
			Where(user.Name(name)).
			Only(ctx)
		if err != nil {
			return err
		}

		mes := &user_stream.UserStreamV1{
			Action:         user_stream.Action_ACTION_MODIFIED,
			PublicId:       updatedUser.UUID.String(),
			Name:           updatedUser.Name,
			Role:           updatedUser.Edges.Role.Name,
			IdempotencyKey: uuid.New().String(),
			Timestamp:      time.Now().UnixNano(),
		}

		_, err = schema.ValidateUserStreamV1(mes)
		if err != nil {
			return err
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

	return updatedUser, err
}
