package api

import (
	"net/http"
)

// InfoResponse object create.
type InfoResponse struct {
	Coins       int64       `json:"coins"`
	Inventory   []Inventory `json:"inventory"`
	CoinHistory CoinHistory `json:"coinHistory"`
}

// CoinHistory object create.
type CoinHistory struct {
	Received []ReceivedTransaction `json:"received"`
	Sent     []SentTransaction     `json:"sent"`
}

// Inventory class create.
type Inventory struct {
	Type     string `json:"type"`
	Quantity int64  `json:"quantity"`
}

// ReceivedTransaction object create.
type ReceivedTransaction struct {
	FromUser string `json:"fromUser"`
	Amount   int64  `json:"amount"`
}

// SentTransaction object create.
type SentTransaction struct {
	ToUser string `json:"toUser"`
	Amount int64  `json:"amount"`
}

// Info handler create.
func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	// userID, ok := middleware.UserIDFromContext(ctx)
	// if !ok {
	// 	h.logger.Println(ok)
	// 	response.Error(w, http.StatusInternalServerError, "not user ID")
	// 	return
	// }

	// info, err := h.service.Info(
	// 	ctx,
	// 	userID,
	// )
	// switch {
	// case errors.Is(err, service.ErrNegativeCoins):

	// }

	w.WriteHeader(http.StatusOK)
}
