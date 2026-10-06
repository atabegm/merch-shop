package purchases

import (
	"context"
	"fmt"
)

// Create create.
func (r *Repo) Create(ctx context.Context, userID, merchID int64) error {
	conn := r.getter.DefaultTrOrDB(
		ctx,
		r.pool,
	)

	_, err := conn.Exec(
		ctx,
		`
		INSERT INTO purchases (user_id, merch_id) 
		VALUES ($1, $2)	
		`,
		userID,
		merchID,
	)
	if err != nil {
		return fmt.Errorf("create purch: %w", err)
	}

	return nil
}
