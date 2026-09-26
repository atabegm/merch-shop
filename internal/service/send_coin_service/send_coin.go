package sendcoinservice

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrInvalidAmount   = errors.New("amount must be positive")
	ErrSelfTrans       = errors.New("cannot send to yourself")
	ErrWithEnoughCoins = errors.New("not enough coins")
)

func (s *service) Send(ctx context.Context, senderID int64, toUser string, amount int64) error {
	if amount < int64(0) {
		return ErrInvalidAmount
	}

	sender, err := s.UserRepository.GetByID(ctx, senderID)
	if err != nil {
		return fmt.Errorf("error with get sender: %w", err)
	}

	receiver, err := s.UserRepository.GetByUsername(ctx, toUser)
	if err != nil {
		return fmt.Errorf("error with get receiver: %w", err)
	}

	if sender.ID == receiver.ID {
		return ErrSelfTrans
	}

	if err := s.CoinTransfersRepository.Send(ctx, senderID, receiver.ID, amount); err != nil {
		return fmt.Errorf("error with send: %w", err)
	}

	return nil
}
