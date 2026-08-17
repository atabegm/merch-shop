package model

// User object create.
type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
}
