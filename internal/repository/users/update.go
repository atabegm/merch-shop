package users

import (
	"context"
	"fmt"
)

// UpdateEmail create.
func (r *Repo) UpdateEmail(ctx context.Context, userID int64, email string) error {
	_, err := r.pool.Exec(ctx, "UPDATE users SET email = $1 WHERE id = $2", email, userID)
	if err != nil {
		return fmt.Errorf("failed to update user. error: %w", err)
	}

	return nil
}
