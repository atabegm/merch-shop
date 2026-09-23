package sendcoinservice

import (
	"avito/internal/model"
	"context"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (model.User, error)
	GetByID(ctx context.Context, id int64) (model.User, error)
}

type CoinTransfersRepository interface {
	Send(ctx context.Context, senderID int64, receiverID int64, amount int64) error
}

type Service struct {
	UserRepository          UserRepository
	CoinTransfersRepository CoinTransfersRepository
}

func New(userReposiroty UserRepository, coinTransfersRepository CoinTransfersRepository) Service {
	return Service{
		UserRepository:          userReposiroty,
		CoinTransfersRepository: coinTransfersRepository,
	}
}
