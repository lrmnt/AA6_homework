package kafka

import "github.com/segmentio/kafka-go"

func NewReader(brokers []string, topic string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  "tasks",
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	})

	return r
}
