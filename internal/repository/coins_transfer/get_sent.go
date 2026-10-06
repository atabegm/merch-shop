package coinstransfer

import (
	"avito/internal/model"
	"context"
	"fmt"
)

// GetSent create.
func (r *Repo) GetSent(ctx context.Context, userID int64) ([]model.SentTransaction, error) {
	rows, err := r.pool.Query(
		ctx,
		`
		SELECT 
			u.username,
			ct.amount
		FROM coins_transfers ct
		JOIN users u ON u.id = ct.receiver_id 
		WHERE ct.sender_id = $1
		`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("error with query sent: %w", err)
	}

	defer rows.Close()
	sentTransactions := make([]model.SentTransaction, 0)

	for rows.Next() {
		var sentTr model.SentTransaction

		err = rows.Scan(
			&sentTr.ToUser,
			&sentTr.Amount,
		)
		if err != nil {
			return nil, fmt.Errorf("error with scan receiver transaction: %w", err)
		}

		sentTransactions = append(sentTransactions, sentTr)
	}

	return sentTransactions, nil
}
