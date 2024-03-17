package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type Producer struct {
	conn *kafka.Conn
}

func NewProducer(ctx context.Context, addr, topic string) (*Producer, error) {
	conn, err := kafka.DialLeader(ctx, "tcp", addr, topic, 0)
	if err != nil {
		return nil, err
	}

	return &Producer{
		conn: conn,
	}, nil
}

func (p *Producer) Produce(message []byte) error {
	_ = p.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

	_, err := p.conn.WriteMessages(kafka.Message{Value: message})

	return err
}

type SaramaProducer struct {
	pr    sarama.SyncProducer
	topic string
}

func NewSaramaProducer(addr string, topic string) *SaramaProducer {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = true
	pr, err := sarama.NewSyncProducer([]string{addr}, cfg)
	if err != nil {
		log.Fatalf("can not create producer: %w", err)
	}

	return &SaramaProducer{
		pr:    pr,
		topic: topic,
	}
}

func (p *SaramaProducer) Send(data []byte) error {
	_, _, err := p.pr.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(data),
	})

	return err
}

func (p *SaramaProducer) Close() error {
	return p.pr.Close()
}
