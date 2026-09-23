package sendcoin

import (
	sendcoinservice "avito/internal/service/send_coin_service"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Service sendcoinservice.Service
	logger  *logrus.Logger
}

func New(service sendcoinservice.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		Service: service,
		logger:  logger,
	}
}
