package service

import (
	"avito/internal/model"
	"context"
	"errors"
	"fmt"
)

var (
	// ErrNegativeCoins create.
	ErrNegativeCoins = errors.New("negative coins")
)

// Info service create.
func (s *Service) Info(ctx context.Context, userID int64) (model.Info, error) {
	coins, err := s.UserRepository.GetCoins(
		ctx,
		userID,
	)
	if err != nil {
		return model.Info{}, fmt.Errorf("get info: %w", err)
	}

	inventory, err := s.PurchasesRepository.GetInventory(
		ctx,
		userID,
	)
	if err != nil {
		return model.Info{}, fmt.Errorf("get info: %w", err)
	}

	received, err := s.CoinTransfersRepository.GetReceived(
		ctx,
		userID,
	)
	if err != nil {
		return model.Info{}, fmt.Errorf("get info: %w", err)
	}

	sent, err := s.CoinTransfersRepository.GetSent(
		ctx,
		userID,
	)
	if err != nil {
		return model.Info{}, fmt.Errorf("get info: %w", err)
	}

	return model.Info{
		Coins:     coins,
		Inventory: inventory,
		CoinHistory: model.CoinHistory{
			Received: received,
			Sent:     sent,
		},
	}, nil
}
