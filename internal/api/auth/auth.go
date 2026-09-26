package auth

import (
	"avito/internal/api/response"
	authservice "avito/internal/service/auth_service"
	"encoding/json"
	"errors"
	"net/http"
)

// Auth handler create.
func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Println(err)
		response.Error(w, http.StatusBadRequest, "error with decode request")
		return
	}

	if err := req.Validate(ctx); err != nil {
		h.logger.Println(err)
		response.Error(w, http.StatusBadRequest, "error with validate")
		return
	}

	token, err := h.service.Auth(ctx, req.Username, req.Password, req.Email)

	switch {
	case errors.Is(err, authservice.ErrInvalidPassword):
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	case err != nil:
		h.logger.Println(err)
		response.Error(w, http.StatusInternalServerError, "server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(AuthResponse{
		Token: token,
	}); err != nil {
		h.logger.Println(err)
		return
	}

}
