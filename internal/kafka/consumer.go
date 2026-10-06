package kafka

import (
	"avito/internal/model"
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Consumer create.
type Consumer struct {
	reader *kafka.Reader
}

// NewConsumer create.
func NewConsumer(broker string, topic string, groupID string) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(
			kafka.ReaderConfig{
				Brokers: []string{broker},
				Topic:   topic,
				GroupID: groupID,
			},
		),
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

		err = json.Unmarshal(
			msg.Value,
			&purchaseEvent,
		)
		if err != nil {
			return fmt.Errorf("error with unmarshal: %w", err)
		}

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
