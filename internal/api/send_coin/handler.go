package sendcoin

import (
	coinstransfer "avito/internal/repository/coins_transfer"
	"avito/internal/repository/users"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	UserRepo          *users.Repo
	CoinsTransferRepo *coinstransfer.Repo
	logger            *logrus.Logger
}

func New(userRepo *users.Repo, coinsTransfer *coinstransfer.Repo, logger *logrus.Logger) Handler {
	return Handler{
		UserRepo:          userRepo,
		CoinsTransferRepo: coinsTransfer,
		logger:            logger,
	}
}
