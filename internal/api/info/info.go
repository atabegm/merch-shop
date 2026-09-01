package info

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response object create.
type Response struct {
	Coins       int64       `json:"coins"`
	CoinHistory CoinHistory `json:"coin_history"`
}

// CoinHistory object create.
type CoinHistory struct {
	Received []ReceivedTransaction `json:"received"`
	Sent     []SentTransaction     `json:"sent"`
}

// ReceivedTransaction object create.
type ReceivedTransaction struct {
	FromUser int64 `json:"from_user"`
	Amount   int64 `json:"amount"`
}

// SentTransaction object create.
type SentTransaction struct {
	ToUser int64 `json:"to_user"`
	Amount int64 `json:"amount"`
}

// Info handler create.
func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// id := ctx.Value("user_id")

	// userID, ok := id.(int)
	// if !ok {
	// 	h.logger.Println("error with converte")
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }
	userID := 1

	usr, err := h.UserRepo.GetByID(ctx, userID)
	if err != nil {
		h.logger.Println("error with user repo", err)
	}

	fmt.Printf("%+v\n", usr)

	coins := usr.Coins

	transfers, err := h.CoinsTransferRepo.GetByUserID(ctx, userID)
	if err != nil {
		h.logger.Println("error with coins transfer repo", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	receivedTransfer := make([]ReceivedTransaction, 0)
	senderTransfer := make([]SentTransaction, 0)

	for _, tr := range transfers {
		if tr.SenderID == usr.ID {
			senderTr := SentTransaction{
				ToUser: tr.ReceiverID,
				Amount: tr.Amount,
			}

			senderTransfer = append(senderTransfer, senderTr)
		}

		if tr.ReceiverID == usr.ID {
			receivedTr := ReceivedTransaction{
				FromUser: tr.SenderID,
				Amount:   tr.Amount,
			}
			receivedTransfer = append(receivedTransfer, receivedTr)
		}
	}

	var coinHistory CoinHistory

	coinHistory.Received = receivedTransfer
	coinHistory.Sent = senderTransfer

	err = json.NewEncoder(w).Encode(Response{
		CoinHistory: coinHistory,
		Coins:       coins,
	})
	if err != nil {
		h.logger.Println("fail to encode. error", err)
	}
}
