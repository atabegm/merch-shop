package info

import (
	coinstransfer "avito/internal/repository/coins_transfer"
	"avito/internal/repository/users"

	"github.com/sirupsen/logrus"
)

// Handler info object create.
type Handler struct {
	CoinsTransferRepo *coinstransfer.Repo
	UserRepo          *users.Repo
	logger            *logrus.Logger
}

// New constructor for user repo.
func New(userRepo *users.Repo, coinsTransferRepo *coinstransfer.Repo, logger *logrus.Logger) *Handler {
	return &Handler{
		UserRepo:          userRepo,
		CoinsTransferRepo: coinsTransferRepo,
		logger:            logger,
	}
}
