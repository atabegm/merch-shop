package consumer

import (
	"avito/internal/apiserver"
	"avito/internal/kafka"
	"avito/internal/service"
	"context"
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	data, err := os.ReadFile(
		"configs/config.yaml",
	)
	if err != nil {
		log.Printf("failed to read yaml file: %v", err)
	}

	cfg := apiserver.NewConfig()

	err = yaml.Unmarshal(
		data,
		cfg,
	)
	if err != nil {
		log.Printf("failed to unmarshal: %v", err)
	}

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
		context.Background(),
		notificationService.HandlePurchase,
	)
	log.Printf("error with consume: %v", err)
}
