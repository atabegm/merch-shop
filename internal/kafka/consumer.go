package kafka

import (
	"avito/internal/model"
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Consumer create.
type Consumer struct {
	reader *kafka.Reader
}

// NewConsumer create.
func NewConsumer(topic, kafkaURL, groupID string) *Consumer {
	return &Consumer{
		kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{kafkaURL},
			GroupID: groupID,
			Topic:   topic,
		}),
	}
}

// Consume create.
func (c *Consumer) Consume(ctx context.Context, handler func(context.Context, *model.PurchaseCreated) error) error {
	for {
		msg, err := c.reader.FetchMessage(
			ctx,
		)
		if err != nil {
			return fmt.Errorf("consumer: %w", err)
		}

		var purchaseEvent model.PurchaseCreated

		err = handler(
			ctx,
			&purchaseEvent,
		)
		if err != nil {
			return fmt.Errorf("consumer: %w", err)
		}

		err = c.reader.CommitMessages(
			ctx,
			msg,
		)
		if err != nil {
			return fmt.Errorf("consumer: %w", err)
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
