package kafka

import (
	"avito/internal/model"
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Producer create.
type Producer struct {
	writer *kafka.Writer
}

// NewProducer create.
func NewProducer(kafkaURL, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(kafkaURL),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

// WriteToKafka create.
func (p *Producer) Produce(ctx context.Context, purchEvent *model.PurchaseCreated) error {
	data, err := json.Marshal(
		purchEvent,
	)
	if err != nil {
		return fmt.Errorf("producer: %w", err)
	}

	err = p.writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   []byte(purchEvent.EventID),
			Value: data,
		},
	)
	if err != nil {
		return fmt.Errorf("producer: %w", err)
	}

	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
