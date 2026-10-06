package merch

import (
	"avito/internal/model"
	"context"
	"fmt"
)

// GetByItem create.
func (r *Repo) GetByName(ctx context.Context, item string) (model.Merch, error) {
	row := r.pool.QueryRow(ctx, "SELECT id, name, price FROM merch WHERE name = $1", item)

	var merch model.Merch

	if err := row.Scan(&merch.ID, &merch.Name, &merch.Price); err != nil {
		return model.Merch{}, fmt.Errorf("error with scan merch: %w", err)
	}

	return merch, nil
}
