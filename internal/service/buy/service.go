package buyservice

import (
	"avito/internal/model"
	"context"
)

type MerchRepository interface {
	GetByName(ctx context.Context, name string) (model.Merch, error)
}

type PurchasesRepository interface {
	Buy(ctx context.Context, userID int64, merchID int64, price int64) error
}

type Service struct {
	PurchasesRepository PurchasesRepository
	MerchRepository     MerchRepository
}

func New(PurchasesRepository PurchasesRepository, MerchRepository MerchRepository) Service {
	return Service{
		PurchasesRepository: PurchasesRepository,
		MerchRepository:     MerchRepository,
	}
}
