package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
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
	_ = p.conn.SetWriteDeadline(time.Now().Add(2 * time.Second))

	_, err := p.conn.WriteMessages(kafka.Message{Value: message})

	return err
}
