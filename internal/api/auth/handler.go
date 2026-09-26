package auth

import (
	"context"

	"github.com/sirupsen/logrus"
)

type AuthService interface {
	Auth(ctx context.Context, username string, password string, email string) (string, error)
}

// Handler for auth.
type Handler struct {
	service   AuthService
	logger    *logrus.Logger
	jwtSecret []byte
}

// New handler for auth.
func New(service AuthService, logger *logrus.Logger, jwtSecret string) *Handler {
	return &Handler{
		service:   service,
		logger:    logger,
		jwtSecret: []byte(jwtSecret),
	}
}