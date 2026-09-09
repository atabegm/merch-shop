package info

import (
	coinstransfer "avito/internal/repository/coins_transfer"
	"avito/internal/repository/merch"
	"avito/internal/repository/purchases"
	"avito/internal/repository/users"

	"github.com/sirupsen/logrus"
)

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
	logger *logrus.Logger) *Handler {

	return &Handler{
		UserRepo:          userRepo,
		CoinsTransferRepo: coinsTransferRepo,
		PurchasesRepo:     purchaseRepo,
		MerchRepo:         merchRepo,
		logger:            logger,
	}
}
