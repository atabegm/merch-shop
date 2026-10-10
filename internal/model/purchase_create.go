package model

import "time"

// PurchaseCreated event.
type PurchaseCreated struct {
	EventID     string    `json:"event_id"`
	UserID      int64     `json:"user_id"`
	Email       string    `json:"email"`
	Item        string    `json:"item"`
	Price       int64     `json:"price"`
	PurchasedAt time.Time `json:"purchased_at"`
}
