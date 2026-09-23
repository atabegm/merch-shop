package info

import (
	"avito/internal/model"
	coinstransfer "avito/internal/repository/coins_transfer"
	"avito/internal/repository/merch"
	"avito/internal/repository/purchases"
	"avito/internal/repository/users"
	"context"

	"github.com/sirupsen/logrus"
)

type CoinTransfersRepository interface {
	GetByUserID(ctx context.Context, id int64) ([]model.Transaction, error)
}

type PurchasesRepository interface {
	GetByUserID(ctx context.Context, userID int64) ([]model.Purchases, error)
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (model.User, error)
}

type MerchRepository interface {
	GetByID(ctx context.Context, id int64) (model.User, error)
}

// Handler info object create.
type Handler struct {
	CoinsTransferRepo *coinstransfer.Repo
	UserRepo          *users.Repo
	PurchasesRepo     *purchases.Repo
	MerchRepo         *merch.Repo
	logger            *logrus.Logger
}

// New constructor for user repo.
func New(userRepo *users.Repo, coinsTransferRepo *coinstransfer.Repo, purchaseRepo *purchases.Repo, merchRepo *merch.Repo,
	logger *logrus.Logger) Handler {

	return Handler{
		UserRepo:          userRepo,
		CoinsTransferRepo: coinsTransferRepo,
		PurchasesRepo:     purchaseRepo,
		MerchRepo:         merchRepo,
		logger:            logger,
	}
}
