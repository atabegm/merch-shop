package coinstransfer

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrWithSend = errors.New("error with send")
)

func (r *Repo) Send(ctx context.Context, senderID int64, receiverID int64, amount int64) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		r.logger.Println(err)
		return fmt.Errorf("error with start transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	res, err := tx.Exec(ctx, "UPDATE users SET coins = coins - $1 WHERE id = $2 AND coins >= $1", amount, senderID)
	if err != nil {
		r.logger.Println(err)
		return fmt.Errorf("error with send coins: %w", err)
	}

	if res.RowsAffected() == 0 {
		return ErrWithSend
	}

	res, err = tx.Exec(ctx, "UPDATE users SET coins = coins + $1 WHERE id = $2", amount, receiverID)
	if err != nil {
		r.logger.Println(err)
		return fmt.Errorf("error with received coins: %w", err)
	}

	if res.RowsAffected() == 0 {
		r.logger.Println(err)
		return fmt.Errorf("TODO: %w", err)
	}

	_, err = tx.Exec(ctx, "INSERT INTO coins_transfers (sender_id, receiver_id, amount) VALUES ($1, $2, $3)", senderID, receiverID, amount)
	if err != nil {
		r.logger.Println(err)
		return fmt.Errorf("error with insert into table coins_transfers: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
