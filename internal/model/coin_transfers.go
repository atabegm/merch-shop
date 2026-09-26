package model

// Transaction object create.
type Transaction struct {
	ID         int64 `json:"id"`
	ReceiverID int64 `json:"receiver_id"`
	SenderID   int64 `json:"sender_id"`
	Amount     int64 `json:"amount"`
}
