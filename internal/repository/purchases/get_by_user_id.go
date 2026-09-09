package purchases

import (
	"avito/internal/model"
	"context"
	"fmt"
)

func (r *Repo) GetByUserID(ctx context.Context, userID int64) ([]model.Purchase, error) {
	rows, err := r.pool.Query(
		ctx,
		"SELECT id, user_id, merch_id, price FROM purchases WHERE user_id = $1",
		userID,
	)
	if err != nil {
		return []model.Purchase{}, fmt.Errorf("error with query in purch repo: %w", err)
	}

	defer rows.Close()

	var purch model.Purchase
	var purchs []model.Purchase

	for rows.Next() {
		if err := rows.Scan(&purch.ID, &purch.UserID, &purch.MerchID, &purch.Price); err != nil {
			return []model.Purchase{}, fmt.Errorf("error with scan purch: %w", err)
		}

		purchs = append(purchs, purch)
	}

	return purchs, nil
}
