package buyservice

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrWithItemName    = errors.New("item name is required")
	ErrWithEnoughCoins = errors.New("not enough coins")
)

func (s *service) Buy(ctx context.Context, userId int64, itemName string) error {
	if itemName == "" {
		return ErrWithItemName
	}

	item, err := s.MerchRepository.GetByName(ctx, itemName)
	if err != nil {
		return fmt.Errorf("error with get item: %w", err)
	}

	user, err := s.UserRepository.GetByID(ctx, userId)
	if err != nil {
		return fmt.Errorf("error with get user by id: %w", err)
	}

	if item.Price > user.Coins {
		return ErrWithEnoughCoins
	}

	err = s.PurchasesRepository.Buy(ctx, userId, item.ID, item.Price)
	if err != nil {
		return err
	}

	return nil
}
