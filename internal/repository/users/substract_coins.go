package users

import (
	"context"
	"fmt"
)

// SubstractCoins create.
func (r *Repo) SubstractCoins(ctx context.Context, amount, userID int64) error {
	conn := r.getter.DefaultTrOrDB(
		ctx,
		r.pool,
	)
	res, err := conn.Exec(
		ctx,
		"UPDATE users SET coins = coins - $1 WHERE id = $2",
		amount,
		userID,
	)
	if err != nil {
		r.logger.Println(err)
		return fmt.Errorf("error with sent coins: %w", err)
	}

	if res.RowsAffected() == 0 {
		return ErrUserNotFound
	}

	return nil
}
