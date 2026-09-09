package sendcoin

type Request struct {
	ToUser string `json:"toUser"`
	Amount int64  `json:"amount"`
}
