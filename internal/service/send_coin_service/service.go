package sendcoinservice

import (
	"avito/internal/model"
	"context"
)

// UserRepository interface create.
type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (model.User, error)
	GetByID(ctx context.Context, id int64) (model.User, error)
}

// CoinTransfersRepository interface create.
type CoinTransfersRepository interface {
	Send(ctx context.Context, senderID int64, receiverID int64, amount int64) error
}

type service struct {
	UserRepository          UserRepository
	CoinTransfersRepository CoinTransfersRepository
}

// New constructor for service sendCoin create.
func New(userReposiroty UserRepository, coinTransfersRepository CoinTransfersRepository) *service {
	return &service{
		UserRepository:          userReposiroty,
		CoinTransfersRepository: coinTransfersRepository,
	}
}
