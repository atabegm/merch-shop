package service_test

import (
	"avito/internal/model"
	"avito/internal/service"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Notification_Integration(t *testing.T) {
	ctx := context.Background()

	notification := service.NewNotificationService(
		"localhost",
		"1025",
		"shop@test.mail.ru",
	)

	event := model.PurchaseCreated{
		Email: "muhammad@mail.ru",
		Item:  "book",
		Price: 50,
	}

	err := notification.HandlePurchase(
		ctx,
		&event,
	)

	assert.NoError(t, err)
}
