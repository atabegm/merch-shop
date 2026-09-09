package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

// Auth handler create.
func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error with decode", http.StatusBadRequest)
		return
	}

	if err := req.Validate(ctx); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.GetByUsername(ctx, req.Username)

	if errors.Is(err, pgx.ErrNoRows) {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "error with hash", http.StatusInternalServerError)
			return
		}

		user, err = h.UserRepo.Create(ctx, req.Username, req.Email, string(hashPassword))
		if err != nil {
			http.Error(w, "error with create", http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		http.Error(w, "error with getting user", http.StatusInternalServerError)
		return
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(req.Password)); err != nil {
			http.Error(w, "wrong password or username", http.StatusUnauthorized)
			return
		}

		if user.Email != req.Email {
			if err := h.UserRepo.UpdateEmail(ctx, user.ID, req.Email); err != nil {
				http.Error(w, "error with update email", http.StatusInternalServerError)
				return
			}

			user.Email = req.Email
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenSigned, err := token.SignedString(h.jwtSecret)
	if err != nil {
		http.Error(w, "error with signed token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&AuthResponse{
		Token: tokenSigned,
	}); err != nil {
		http.Error(w, "error with response", http.StatusInternalServerError)
		return
	}
}
