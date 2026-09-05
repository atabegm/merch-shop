package users

import (
	"avito/internal/model"
	"context"
	"fmt"
)

// Create user create.
func (r *Repo) Create(ctx context.Context, username, email, hashPassword string) (model.User, error) {
	row := r.pool.QueryRow(
		ctx,
		"INSERT INTO users (username, email, hash_password) VALUES ($1, $2, $3) RETURNING id,username,email,hash_password,coins",
		username,
		email,
		hashPassword,
	)

	var usr model.User

	err := row.Scan(
		&usr.ID,
		&usr.Username,
		&usr.Email,
		&usr.HashPassword,
		&usr.Coins,
	)
	if err != nil {
		return model.User{}, fmt.Errorf("error with create user: %w", err)
	}

	return usr, nil
}
