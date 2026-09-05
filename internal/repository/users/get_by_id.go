package users

import (
	"avito/internal/model"
	"context"
	"fmt"
)

// GetByID TODO.
func (r *Repo) GetByID(ctx context.Context, id int) (model.User, error) {
	row := r.pool.QueryRow(ctx, "SELECT id, username, email, hash_password, coins FROM users WHERE id = $1", id)

	var usr model.User
	err := row.Scan(&usr.ID, &usr.Username, &usr.Email, &usr.HashPassword, &usr.Coins)
	if err != nil {
		return model.User{}, fmt.Errorf("user repo: %w", err)
	}

	return usr, nil
}
