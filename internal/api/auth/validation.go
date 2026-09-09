package auth

import (
	"context"
	"fmt"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

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
