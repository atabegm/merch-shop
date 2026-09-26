package buyservice

import (
	"avito/internal/model"
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (model.User, error)
}

type MerchRepository interface {
	GetByName(ctx context.Context, name string) (model.Merch, error)
}

type PurchasesRepository interface {
	Buy(ctx context.Context, userID int64, merchID int64, price int64) error
}

type service struct {
	PurchasesRepository PurchasesRepository
	MerchRepository     MerchRepository
	UserRepository      UserRepository
}

func New(PurchasesRepository PurchasesRepository, MerchRepository MerchRepository, UserRepository UserRepository) *service {
	return &service{
		PurchasesRepository: PurchasesRepository,
		MerchRepository:     MerchRepository,
		UserRepository:      UserRepository,
	}
}
