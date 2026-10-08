package api

import (
	"avito/internal/api/middleware"
	"avito/internal/api/response"
	"avito/internal/model"
	"encoding/json"
	"net/http"
)

// InfoResponse object create.
type InfoResponse struct {
	Coins       int64                 `json:"coins"`
	Inventory   []model.InventoryItem `json:"inventory"`
	CoinHistory model.CoinHistory     `json:"coinHistory"`
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

	info, err := h.service.Info(
		ctx,
		userID,
	)
	if err != nil {
		h.logger.Println(err)
		response.Error(w, http.StatusInternalServerError, "error with info")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(InfoResponse{
		Coins:       info.Coins,
		Inventory:   info.Inventory,
		CoinHistory: info.CoinHistory,
	})
	if err != nil {
		h.logger.Println(err)
	}
}
