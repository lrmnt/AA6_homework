package service

import (
	"context"
	"fmt"
	"github.com/lrmnt/AA6_homework/auth/ent"
	"github.com/lrmnt/AA6_homework/lib/kafka"
	"go.uber.org/zap"
)

type Service struct {
	log          *zap.Logger
	client       *ent.Client
	userProducer *kafka.Producer
}

func New(log *zap.Logger, client *ent.Client, userProducer *kafka.Producer) *Service {
	return &Service{
		log:          log,
		client:       client,
		userProducer: userProducer,
	}
}

func (s *Service) tx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("can not start tx: %w", err)
	}

	err = fn(tx)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("can not rollback tx: %w, Err: %w", rbErr, err)
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("can not commit tx: %w", err)
	}

	return nil
}
