package sendcoin

import (
	"context"

	"github.com/sirupsen/logrus"
)

type SendCoinService interface {
	Send(ctx context.Context, senderID int64, toUser string, amount int64) error
}
type Handler struct {
	service SendCoinService
	logger  *logrus.Logger
}

func New(service SendCoinService, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
