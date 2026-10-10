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
			m.name,
			COUNT(*)
		FROM purchases p
		JOIN merch m ON m.id = p.merch_id
		WHERE p.user_id = $1
		GROUP BY(m.name)
		`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("get inv: %w", err)
	}

	defer rows.Close()

	var inventory []model.InventoryItem
	var inventoryItem model.InventoryItem

	for rows.Next() {
		err = rows.Scan(
			&inventoryItem.Type,
			&inventoryItem.Quantity,
		)
		if err != nil {
			return nil, fmt.Errorf("scan inv item: %w", err)
		}

		inventory = append(inventory, inventoryItem)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows err inv item: %w", err)
	}

	return inventory, nil
}
