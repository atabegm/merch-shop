package api

import (
	"avito/internal/api/middleware"
	"avito/internal/api/response"
	"avito/internal/service"
	"errors"
	"net/http"
)

// Buy handler create.
func (h *Handler) Buy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		h.logger.Println("error with middleware in Buy api")
		response.Error(w, http.StatusInternalServerError, "error with user ID")
		return
	}

	itemName := r.PathValue("item")

	err := h.service.Buy(ctx, userID, itemName)

	switch {
	case errors.Is(err, service.ErrEmptyItem):
		h.logger.Printf("buy error: %v", err)
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	case errors.Is(err, service.ErrNotEnoughCoins):
		h.logger.Printf("buy error: %v", err)
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	case err != nil:
		h.logger.Printf("buy error: %v", err)
		response.Error(w, http.StatusInternalServerError, "error with buy")
		return
	}

	w.WriteHeader(http.StatusOK)
}
