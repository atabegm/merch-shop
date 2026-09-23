package model

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestUser(t *testing.T) *User {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	return &User{
		Username:     "muhammadrasul",
		Email:        "muhammadrasul@mail.ru",
		HashPassword: string(hash),
	}
}
