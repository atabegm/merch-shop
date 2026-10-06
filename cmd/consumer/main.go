package main

import (
	"avito/internal/apiserver"
	"avito/internal/kafka"
	"avito/internal/service"
	"context"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

func main() {
	cfg := apiserver.NewConfig()

	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		log.Fatal("failed to read config file:", err)
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Fatal("failed to unmarshal config:", err)
	}

	ctx := context.Background()

	consumer := kafka.NewConsumer(
		cfg.KafkaBroker,
		cfg.KafkaTopic,
		cfg.KafkaGroupID,
	)
	defer consumer.Close()

	notificationService := service.NewNotificationService(
		cfg.SMTPHost,
		cfg.SMTPPort,
		cfg.SMTPFrom,
	)

	err = consumer.Consume(
		ctx,
		notificationService.HandlePurchase,
	)

	log.Fatal("consumer stopped:", err)
}
