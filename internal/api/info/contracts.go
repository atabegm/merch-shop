package info

import (
	"avito/internal/model"
	"context"
)

// CoinTransfersRepository for info handler create.
type CoinTransfersRepository interface {
	GetByUserID(ctx context.Context, id int64) ([]model.Transaction, error)
}

// PurchasesRepository for info handler create.
type PurchasesRepository interface {
	GetByUserID(ctx context.Context, userID int64) ([]model.Purchases, error)
}

// UserRepository for info handler create.
type UserRepository interface {
	GetByID(ctx context.Context, id int64) (model.User, error)
}

// MerchRepository for info handler create.
type MerchRepository interface {
	GetByID(ctx context.Context, id int64) (model.Merch, error)
}
