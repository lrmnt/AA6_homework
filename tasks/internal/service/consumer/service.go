package consumer

import (
	"github.com/lrmnt/AA6_homework/tasks/ent"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Service struct {
	log                  *zap.Logger
	client               *ent.Client
	userStreamV1Consumer *kafka.Reader
}

func New(l *zap.Logger, client *ent.Client,
	userStreamV1Consumer *kafka.Reader,
) *Service {
	return &Service{
		log:                  l,
		client:               client,
		userStreamV1Consumer: userStreamV1Consumer,
	}
}
