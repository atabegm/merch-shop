package auth

import (
	authservice "avito/internal/service/auth_service"

	"github.com/sirupsen/logrus"
)

// Handler for auth.
type Handler struct {
	UserRepo  authservice.Service
	logger    *logrus.Logger
	jwtSecret []byte
}


// New handler for auth.
func New(service authservice.Service, logger *logrus.Logger, jwtSecret string) *Handler {
	return &Handler{
		UserRepo:  service,
		logger:    logger,
		jwtSecret: []byte(jwtSecret),
	}
}
