package kafka

import (
	"github.com/IBM/sarama"
	"github.com/segmentio/kafka-go"
	"log"
)

func NewReader(brokers []string, topic, group string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  group,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	})

	return r
}

func NewSaramaReader(brokers []string, group string) sarama.ConsumerGroup {
	conf := sarama.NewConfig()
	conf.Consumer.Return.Errors = true

	cg, err := sarama.NewConsumerGroup(brokers, group, conf)
	if err != nil {
		log.Fatalf("can not create consumer group: %w", err)
	}

	return cg
}
