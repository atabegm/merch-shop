package sendcoin

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
)

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req Request

	senderID := int64(1)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error with decode requests body", http.StatusBadRequest)
		return
	}

	if err := validation.ValidateStruct(
		&req,
		validation.Field(&req.ToUser, validation.Required), 	
		validation.Field(&req.Amount, validation.Min(int64(1))),
	); err != nil {
		http.Error(w, "error with request", http.StatusBadRequest)
		return
	}

	receiver, err := h.UserRepo.GetByUsername(ctx, req.ToUser)
	if err != nil {
		http.Error(w, "error with get receiver", http.StatusInternalServerError)
		return
	}

	sender, err := h.UserRepo.GetByID(ctx, senderID)
	if err != nil {
		http.Error(w, "error with get sender", http.StatusInternalServerError)
		return
	}

	if sender.ID == receiver.ID {
		http.Error(w, "Stop! You cant send coins to yourself", http.StatusBadRequest)
		return
	}

	if req.Amount > sender.Coins {
		http.Error(w, "Error! Amount less then coins u have", http.StatusBadRequest)
		return
	}

	if err := h.CoinsTransferRepo.Send(ctx, sender.ID, receiver.ID, req.Amount); err != nil {
		http.Error(w, "error with send", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
