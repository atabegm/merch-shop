package model

// User object create.
type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
	Coins        int64  `json:"coins"`
}
