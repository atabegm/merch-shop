package buy

import (
	"avito/internal/api/auth/middleware"
	"avito/internal/api/response"
	"avito/internal/repository/purchases"
	buyservice "avito/internal/service/buy"
	"errors"
	"net/http"
)

func (h *Handler) Buy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		h.logger.Println("error with middleware in Buy api")
		response.Error(w, http.StatusInternalServerError, "error with user ID")
		return
	}

	itemName := r.PathValue("item")

	err := h.Service.Buy(ctx, userID, itemName)

	switch {
	case errors.Is(err, buyservice.ErrWithItemName):
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	case errors.Is(err, purchases.ErrNotEnoughCoins):
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	case err != nil:
		response.Error(w, http.StatusInternalServerError, "error with buy")
		return
	}

	w.WriteHeader(http.StatusOK)
}
