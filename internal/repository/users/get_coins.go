package users

import (
	"context"
	"fmt"
)

// GetCoins create.
func (r *Repo) GetCoins(ctx context.Context, userID int64) (int64, error) {
	var coins int64

	row := r.pool.QueryRow(
		ctx,
		"SELECT coins FROM users WHERE id = $1",
		userID,
	)

	err := row.Scan(&coins)
	if err != nil {
		return 0, fmt.Errorf("error with scan: %w", err)
	}

	return coins, nil
}
