package sendcoin

import (
	"avito/internal/api/auth/middleware"
	"avito/internal/api/response"
	sendcoinservice "avito/internal/service/send_coin_service"
	"encoding/json"
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
)

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req Request

	senderID, ok := middleware.UserIDFromContext(ctx)
	if !ok {
		h.logger.Println("error in send handler", ok)
		response.Error(w, http.StatusInternalServerError, "error with sender ID")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "error with decode request's body")
		return
	}

	if err := validation.ValidateStruct(
		&req,
		validation.Field(&req.ToUser, validation.Required),
		validation.Field(&req.Amount, validation.Required),
	); err != nil {
		h.logger.Println(err)
		response.Error(w, http.StatusBadRequest, "error with request")
		return
	}

	err := h.service.Send(ctx, senderID, req.ToUser, req.Amount)

	switch {
	case errors.Is(err, sendcoinservice.ErrInvalidAmount):
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	case errors.Is(err, sendcoinservice.ErrSelfTrans):
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	case errors.Is(err, sendcoinservice.ErrWithEnoughCoins):
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	case err != nil:
		response.Error(w, http.StatusInternalServerError, "error with send")
		return
	}

	w.WriteHeader(http.StatusOK)
}
