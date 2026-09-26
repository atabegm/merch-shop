package buy

import (
	"context"

	"github.com/sirupsen/logrus"
)

type BuyService interface {
	Buy(ctx context.Context, userId int64, itemName string) error
}

type Handler struct {
	service BuyService
	logger  *logrus.Logger
}

func New(service BuyService, logger *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
