package service

import (
	"avito/internal/model"
	"context"
)

// UserRepository create.
type UserRepository interface {
	AddCoins(ctx context.Context, amount, userID int64) error
	SubstractCoins(ctx context.Context, amount, userID int64) error
	GetCoins(ctx context.Context, userID int64) (int64, error)
	GetByID(ctx context.Context, id int64) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	Create(ctx context.Context, username, email, hashPassword string) (model.User, error)
	UpdateEmail(ctx context.Context, userID int64, email string) error
}

// CoinTransfersRepository for info handler create.
type CoinTransfersRepository interface {
	GetReceived(ctx context.Context, userID int64) ([]model.ReceivedTransaction, error)
	GetSent(ctx context.Context, userID int64) ([]model.SentTransaction, error)
}

// PurchasesRepository for info handler create.
type PurchasesRepository interface {
	GetInventory(ctx context.Context, userID int64) ([]model.InventoryItem, error)
	Create(ctx context.Context, userID, merchID int64) error
}

// MerchRepository create.
type MerchRepository interface {
	GetByName(ctx context.Context, name string) (model.Merch, error)
}

// Transactor create.
type Transactor interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}

// Produce interface create.
type Produce interface {
	Produce(ctx context.Context, purchEvent *model.PurchaseCreated) error
}
