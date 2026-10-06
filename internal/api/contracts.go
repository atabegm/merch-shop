package api

import (
	"avito/internal/model"
	"context"
)

// Service create.
type Service interface {
	Buy(ctx context.Context, userID int64, itemName string) error

	SendCoins(ctx context.Context, senderID int64, toUser string, amount int64) error

	Auth(ctx context.Context, username string, password string, email string) (string, error)

	Info(ctx context.Context, userID int64) (model.Info, error)
}
