package service

import (
	"avito/internal/model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (

	// ErrEmptyItem create.
	ErrEmptyItem = errors.New("item name is required")
	// ErrNotEnoughCoins create.
	ErrNotEnoughCoins = errors.New("not enough coins")
)

// Buy create.
func (s *Service) Buy(ctx context.Context, userID int64, itemName string) error {
	if itemName == "" {
		return ErrEmptyItem
	}

	merch, err := s.MerchRepository.GetByName(
		ctx,
		itemName,
	)
	if err != nil {
		return fmt.Errorf("buy service: %w", err)
	}

	user, err := s.UserRepository.GetByID(
		ctx,
		userID,
	)
	if err != nil {
		return fmt.Errorf("buy service: %w", err)
	}

	err = s.Transactor.Do(
		ctx,
		func(ctx context.Context) error {
			err = s.PurchasesRepository.Create(
				ctx,
				userID,
				merch.ID,
			)
			if err != nil {
				return fmt.Errorf("buy service: %w", err)
			}

			err = s.UserRepository.SubstractCoins(
				ctx,
				merch.Price,
				userID,
			)
			if err != nil {
				return fmt.Errorf("buy service: %w", err)
			}

			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("buy service: %w", err)
	}

	purchEvent := model.PurchaseCreated{
		EventID:     uuid.NewString(),
		UserID:      userID,
		Email:       user.Email,
		Item:        merch.Name,
		Price:       merch.Price,
		PurchasedAt: time.Now(),
	}

	err = s.Produce.Produce(
		ctx,
		&purchEvent,
	)
	if err != nil {
		return fmt.Errorf("buy service: %w", err)
	}

	return nil
}
