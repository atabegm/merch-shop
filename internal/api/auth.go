package api

import (
	"avito/internal/api/response"
	"avito/internal/service"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Request object create.
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Response object for auth create.
type AuthResponse struct {
	Token string `json:"token"`
}

// Validate func create.
func (r *AuthRequest) Validate(ctx context.Context) error {
	if err := validation.ValidateStructWithContext(
		ctx,
		r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required),
		validation.Field(&r.Username, validation.Required),
	); err != nil {
		return fmt.Errorf("error with validate: %w", err)
	}

	return nil
}

// Auth handler create.
func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AuthRequest

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
	case errors.Is(err, service.ErrInvalidPassword):
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
