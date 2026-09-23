package model

// Purchases object create.
type Purchases struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	MerchID int64 `json:"merch_id"`
	Price   int64 `json:"price"`
}
