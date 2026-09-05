package users

import (
	"avito/internal/model"
	"context"
	"fmt"
)

// GetByUsername TODO.
func (r *Repo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	row := r.pool.QueryRow(ctx, "SELECT id, username, email, hash_password, coins FROM users WHERE username = $1", username)

	var usr model.User
	if err := row.Scan(&usr.ID, &usr.Username, &usr.Email, &usr.HashPassword, &usr.Coins); err != nil {
		return model.User{}, fmt.Errorf("error with username scan: %w", err)
	}

	return usr, nil
}
