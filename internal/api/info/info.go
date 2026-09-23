package info

import (
	"avito/internal/api/auth/middleware"
	"avito/internal/api/response"
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

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		h.logger.Println(ok)
		response.Error(w, http.StatusInternalServerError, "not user ID")
		return
	}

	usr, err := h.UserRepo.GetByID(ctx, userID)
	if err != nil {
		h.logger.Println(err)
		response.Error(w, http.StatusInternalServerError, "error with get user")
		return
	}

	coins := usr.Coins

	transfers, err := h.CoinsTransferRepo.GetByUserID(ctx, userID)
	if err != nil {
		h.logger.Println(err)
		response.Error(w, http.StatusInternalServerError, "server error")
		return
	}

	receivedTransactions := make([]ReceivedTransaction, 0)
	senderTransactions := make([]SentTransaction, 0)

	for _, tr := range transfers {
		if tr.ReceiverID == usr.ID {
			sender, err := h.UserRepo.GetByID(ctx, tr.SenderID)
			if err != nil {
				h.logger.Println(err)
				response.Error(w, http.StatusInternalServerError, "error with get receiver")
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
				h.logger.Println(err)
				response.Error(w, http.StatusInternalServerError, "error with get sender")
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
		h.logger.Println(err)
		response.Error(w, http.StatusInternalServerError, "error with get purchases")
		return
	}

	inventoryMap := make(map[string]int64)

	for _, purch := range purchases {
		merch, err := h.MerchRepo.GetByID(ctx, purch.MerchID)
		if err != nil {
			h.logger.Println(err)
			response.Error(w, http.StatusInternalServerError, "error with get merch")
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
		h.logger.Println(err)
		response.Error(w, http.StatusInternalServerError, "fail to encode info response")
		return
	}
}
