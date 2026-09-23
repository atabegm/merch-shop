package authservice

import (
	"avito/internal/model"
	"context"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (model.User, error)
	Create(ctx context.Context, username, email, hashPassword string) (model.User, error)
	UpdateEmail(ctx context.Context, userID int64, email string) error
}

type Service struct {
	UserRepository UserRepository
	jwtSecret      string
}

func New(UserRepository UserRepository, jwtSecret string) Service {
	return Service{
		UserRepository: UserRepository,
		jwtSecret:      jwtSecret,
	}
}
