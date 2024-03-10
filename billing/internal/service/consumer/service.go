package consumer

import (
	"github.com/lrmnt/AA6_homework/billing/ent"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Service struct {
	log    *zap.Logger
	client *ent.Client

	userStreamV0Consumer *kafka.Reader
	userStreamV1Consumer *kafka.Reader
	taskStreamV0Consumer *kafka.Reader
	taskStreamV1Consumer *kafka.Reader

	taskEventV1Consumer *kafka.Reader
}

func New(
	l *zap.Logger,
	client *ent.Client,
	userStreamV0Consumer *kafka.Reader,
	userStreamV1Consumer *kafka.Reader,
	taskStreamV0Consumer *kafka.Reader,
	taskStreamV1Consumer *kafka.Reader,
	taskEventV1Consumer *kafka.Reader,
) *Service {
	return &Service{
		log:                  l,
		client:               client,
		userStreamV0Consumer: userStreamV0Consumer,
		userStreamV1Consumer: userStreamV1Consumer,
		taskStreamV0Consumer: taskStreamV0Consumer,
		taskStreamV1Consumer: taskStreamV1Consumer,
		taskEventV1Consumer:  taskEventV1Consumer,
	}
}
