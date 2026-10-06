package service

import (
	"context"
	"errors"
	"fmt"
)

var (
	// ErrEmptyAmount create.
	ErrEmptyAmount = errors.New("empty amount")
	// ErrSelfTrans create.
	ErrSelfTrans = errors.New("self transaction")
)

// SendCoins create.
func (s *Service) SendCoins(ctx context.Context, userID int64, toUser string, amount int64) error {
	if amount < 0 {
		return ErrEmptyAmount
	}
	err := s.Transactor.Do(
		ctx,
		func(ctx context.Context) error {
			receiver, err := s.UserRepository.GetByUsername(
				ctx,
				toUser,
			)
			if err != nil {
				return fmt.Errorf("send coins: %w", err)
			}

			if receiver.ID == userID {
				return ErrSelfTrans
			}

			err = s.UserRepository.AddCoins(
				ctx,
				amount,
				receiver.ID,
			)
			if err != nil {
				return fmt.Errorf("send coins: %w", err)
			}

			err = s.UserRepository.SubstractCoins(
				ctx,
				amount,
				userID,
			)
			if err != nil {
				return fmt.Errorf("send coins: %w", err)
			}

			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("send coins: %w", err)
	}

	return nil
}
