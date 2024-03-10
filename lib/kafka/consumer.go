package kafka

import "github.com/segmentio/kafka-go"

func NewReader(brokers []string, topic, group string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  group,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	})

	return r
}
