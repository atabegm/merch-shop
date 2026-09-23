package purchases

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrNotEnoughCoins = errors.New("not enough coins")
)

func (r *Repo) Buy(ctx context.Context, userID int64, merchID int64, price int64) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error with begin transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	res, err := tx.Exec(ctx, "UPDATE users SET coins = coins - $1 WHERE id = $2 AND coins >= $1", price, userID)
	if err != nil {
		return fmt.Errorf("error with update coins: %w", err)
	}

	if res.RowsAffected() == 0 {
		return ErrNotEnoughCoins
	}

	_, err = tx.Exec(ctx, "INSERT INTO purchases (user_id, merch_id, price) VALUES ($1, $2, $3)", userID, merchID, price)
	if err != nil {
		return fmt.Errorf("error with insert into purchse: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error with commit in buy func: %w", err)
	}

	return nil
}
