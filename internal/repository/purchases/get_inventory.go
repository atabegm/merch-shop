package purchases

import (
	"avito/internal/model"
	"context"
	"fmt"
)

// GetInventory create.
func (r *Repo) GetInventory(ctx context.Context, userID int64) ([]model.InventoryItem, error) {
	rows, err := r.pool.Query(
		ctx,
		`
		SELECT 
			item
			COUNT(*)
		FROM purchases
		WHERE user_id = $1
		GROUP BY item
		`,
		userID,
	)
	if err != nil {
		return []model.InventoryItem{}, fmt.Errorf("get inventory: %w", err)
	}

	defer rows.Close()

	var inventory []model.InventoryItem

	for rows.Next() {
		var item model.InventoryItem

		err = rows.Scan(
			&item.Type,
			&item.Quantity,
		)
		if err != nil {
			return nil, fmt.Errorf("error with scan: %w", err)
		}

		inventory = append(inventory, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}

	return inventory, nil
}
