package model

// Info struct for service create.
type Info struct {
	Coins       int64           `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}

// CoinHistory create.
type CoinHistory struct {
	Received []ReceivedTransaction `json:"received"`
	Sent     []SentTransaction     `json:"sent"`
}

// InventoryItem create.
type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int64  `json:"quantity"`
}

// ReceivedTransaction create.
type ReceivedTransaction struct {
	FromUser string `json:"fromUser"`
	Amount   int64  `json:"amount"`
}

// SentTransaction create.
type SentTransaction struct {
	ToUser string `json:"toUser"`
	Amount int64  `json:"amount"`
}
