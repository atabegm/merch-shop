package model

// Purchase object create.
type Purchase struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	MerchID int64 `json:"merch_id"`
	Price   int64 `json:"price"`
}
