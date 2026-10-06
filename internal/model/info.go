package model

type Info struct {
	Coins       int64
	Inventory   []InventoryItem
	CoinHistory CoinHistory
}

type CoinHistory struct {
	Received []ReceivedTransaction
	Sent     []SentTransaction
}

type InventoryItem struct {
	Type     string
	Quantity int64
}

type ReceivedTransaction struct {
	FromUser string
	Amount   int64
}

type SentTransaction struct {
	ToUser string
	Amount int64
}