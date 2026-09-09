package info

import (
	"encoding/json"
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
	ctx := r.Context()
	// id := ctx.Value("user_id")

	// userID, ok := id.(int64)
	// if !ok {
	// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	userID := int64(1)

	usr, err := h.UserRepo.GetByID(ctx, userID)
	if err != nil {
		http.Error(w, "error with get user", http.StatusInternalServerError)
		return
	}

	coins := usr.Coins

	transfers, err := h.CoinsTransferRepo.GetByUserID(ctx, userID)
	if err != nil {
		h.logger.Println("error with coins transfer repo", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	receivedTransactions := make([]ReceivedTransaction, 0)
	senderTransactions := make([]SentTransaction, 0)

	for _, tr := range transfers {
		if tr.ReceiverID == usr.ID {
			sender, err := h.UserRepo.GetByID(ctx, tr.SenderID)
			if err != nil {
				http.Error(w, "error with get receiver", http.StatusInternalServerError)
				return
			}
			receivedTr := ReceivedTransaction{
				FromUser: sender.Username,
				Amount:   tr.Amount,
			}

			receivedTransactions = append(receivedTransactions, receivedTr)
		}

		if tr.SenderID == usr.ID {
			receiver, err := h.UserRepo.GetByID(ctx, tr.ReceiverID)
			if err != nil {
				http.Error(w, "error with get sender", http.StatusInternalServerError)
				return
			}

			senderTr := SentTransaction{
				ToUser: receiver.Username,
				Amount: tr.Amount,
			}

			senderTransactions = append(senderTransactions, senderTr)
		}
	}

	var coinHistory CoinHistory

	coinHistory.Received = receivedTransactions
	coinHistory.Sent = senderTransactions

	purchases, err := h.PurchasesRepo.GetByUserID(ctx, usr.ID)
	if err != nil {
		http.Error(w, "error with get purchases", http.StatusInternalServerError)
		return
	}

	inventoryMap := make(map[string]int64)

	for _, purch := range purchases {
		merch, err := h.MerchRepo.GetByID(ctx, purch.MerchID)
		if err != nil {
			http.Error(w, "error with get merch", http.StatusInternalServerError)
			return
		}

		inventoryMap[merch.Name]++
	}

	inventory := make([]Inventory, 0)

	for name, quantity := range inventoryMap {
		inventory = append(inventory, Inventory{
			Type:     name,
			Quantity: quantity,
		})
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(InfoResponse{
		CoinHistory: coinHistory,
		Inventory:   inventory,
		Coins:       coins,
	})
	if err != nil {
		h.logger.Println("fail to encode. error", err)
	}
}
