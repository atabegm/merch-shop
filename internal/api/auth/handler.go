package auth

import (
	"avito/internal/repository/users"

	"github.com/sirupsen/logrus"
)

// Handler for auth.
type Handler struct {
	UserRepo  *users.Repo
	logger    *logrus.Logger
	jwtSecret []byte
}

// New handler for auth.
func New(userRepo *users.Repo, logger *logrus.Logger, jwtSecret string) *Handler {
	return &Handler{
		UserRepo:  userRepo,
		logger:    logger,
		jwtSecret: []byte(jwtSecret),
	}
}
