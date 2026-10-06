package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrInvalidPassword create.
	ErrInvalidPassword = errors.New("invalid password")
	// ErrEmptyUsername create.
	ErrEmptyUsername = errors.New("empty username")
)

// Auth create.
func (s *Service) Auth(ctx context.Context, username string, password string, email string) (string, error) {
	if password == "" {
		return "", ErrInvalidPassword
	}

	if username == "" {
		return "", ErrEmptyUsername
	}

	user, err := s.UserRepository.GetByUsername(ctx, username)
	if errors.Is(err, pgx.ErrNoRows) {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return "", fmt.Errorf("error with hash password: %w", err)
		}

		user, err = s.UserRepository.Create(ctx, username, email, string(hashPassword))
		if err != nil {
			return "", fmt.Errorf("error with create user: %w", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("error with get user by username: %w", err)
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
		if err != nil {
			return "", ErrInvalidPassword
		}

		if user.Email != email {
			if err := s.UserRepository.UpdateEmail(ctx, user.ID, email); err != nil {
				return "", fmt.Errorf("error with update email: %w", err)
			}
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenSighed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error with sigh token: %w", err)
	}

	return tokenSighed, nil
}
