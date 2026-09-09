package coinstransfer

import (
	"avito/internal/model"
	"context"
	"fmt"
)

// GetByUserID - find transactions.
func (r *Repo) GetByUserID(ctx context.Context, id int64) ([]model.Transaction, error) {
	rows, err := r.pool.Query(
		ctx,
		"SELECT id, receiver_id, sender_id, amount FROM coins_transfers WHERE sender_id = $1 OR receiver_id = $1", id,
	)
	if err != nil {
		return nil, fmt.Errorf("fail with coins transfer. error: %w", err)
	}

	defer rows.Close()

	var tr model.Transaction
	var trs []model.Transaction

	for rows.Next() {
		err = rows.Scan(&tr.ID, &tr.ReceiverID, &tr.SenderID, &tr.Amount)
		if err != nil {
			return nil, err
		}

		trs = append(trs, tr)
	}

	return trs, nil
}
