package coinstransfer

import (
	"avito/internal/model"
	"context"
	"fmt"
)

// GetReceived create.
func (r *Repo) GetReceived(ctx context.Context, userID int64) ([]model.ReceivedTransaction, error) {
	rows, err := r.pool.Query(
		ctx,
		`
		SELECT 
			u.username,
			ct.amount
		FROM coins_transfers ct
		JOIN users u ON u.id = ct.sender_id
		WHERE ct.receiver_id = $1
		`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("error with query received: %w", err)
	}

	defer rows.Close()
	receivedTransactions := make([]model.ReceivedTransaction, 0)

	for rows.Next() {
		var receivedTr model.ReceivedTransaction

		err := rows.Scan(
			&receivedTr.FromUser,
			&receivedTr.Amount,
		)
		if err != nil {
			return nil, fmt.Errorf("error with scan receiver transaction: %w", err)
		}

		receivedTransactions = append(receivedTransactions, receivedTr)
	}

	return receivedTransactions, nil
}
