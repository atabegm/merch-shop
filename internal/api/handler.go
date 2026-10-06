package api

import (
	"github.com/sirupsen/logrus"
)

// Handler for auth.
type Handler struct {
	service Service
	logger  *logrus.Logger
}

// New handler create.
func New(
	service Service,
	logger *logrus.Logger,
) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
