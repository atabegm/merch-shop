package coinstransfer

import (
	"context"
	"fmt"
)

func (r *Repo) Send(ctx context.Context, senderID int64, receiverID int64, amount int64) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error with pool begin: %w", err)
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "UPDATE users SET coins = coins - $1 WHERE id = $2", amount, senderID)
	if err != nil {
		return fmt.Errorf("error with send: %w", err)
	}

	_, err = tx.Exec(ctx, "UPDATE users SET coins = coins + $1 WHERE id = $2", amount, receiverID)
	if err != nil {
		return fmt.Errorf("error with receive: %w", err)
	}
	
	_, err = tx.Exec(ctx,
		"INSERT INTO coins_transfers (sender_id, receiver_id, amount) VALUES ($1, $2, $3)",
		senderID, receiverID, amount,
	)

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error with commit: %w", err)
	}

	return nil
}
