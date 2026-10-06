package service

import (
	"avito/internal/model"
	"context"
	"fmt"
	"net/smtp"
)

type NotificationService struct {
	host string
	port string
	from string
}

// NewNotificationService create.
func NewNotificationService(host, port, from string) *NotificationService {
	return &NotificationService{
		host: host,
		port: port,
		from: from,
	}
}

// HandlePurchase create.
func (s *NotificationService) HandlePurchase(ctx context.Context, purchEvent *model.PurchaseCreated) error {
	addr := s.host + ":" + s.port

	to := []string{
		purchEvent.Email,
	}

	body := fmt.Sprintf(
		"U bought %s for %d coins",
		purchEvent.Item,
		purchEvent.Price,
	)

	message := []byte(
		"From: " + s.from + "\r\n" +
			"To: " + purchEvent.Email + "\r\n" +
			"Subject: Purchase successful\r\n" +
			"\r\n" +
			body + "\r\n",
	)

	fmt.Printf("EMAIL: %q\n", purchEvent.Email)

	err := smtp.SendMail(
		addr,
		nil,
		s.from,
		to,
		message,
	)
	if err != nil {
		return fmt.Errorf("notification: %w", err)
	}

	return nil
}
