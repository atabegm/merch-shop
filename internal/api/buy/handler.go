package buy

import (
	buyservice "avito/internal/service/buy"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Service buyservice.Service
	logger  *logrus.Logger
}

func New(service buyservice.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		Service: service,
		logger:  logger,
	}
}
