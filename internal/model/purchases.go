package model

import "time"

// Purchases object create.
type Purchases struct {
	ID          int64         `json:"id"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Item        string        `json:"item"`
	Price       int64         `json:"price"`
	PurchasedAt time.Duration `json:"purchased_at"`
}
