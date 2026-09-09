package merch

import (
	"avito/internal/model"
	"context"
	"fmt"
)

func (r *Repo) GetByID(ctx context.Context, merchID int64) (model.Merch, error) {
	row := r.pool.QueryRow(ctx, "SELECT id, name, price FROM merch WHERE id = $1", merchID)

	var merch model.Merch

	if err := row.Scan(&merch.ID, &merch.Name, &merch.Price); err != nil {
		return model.Merch{}, fmt.Errorf("error with scan merch: %w", err)
	}

	return merch, nil
}
