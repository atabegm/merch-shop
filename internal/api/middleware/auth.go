package middleware

import (
	"avito/internal/api/response"
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type ctxKey string

var userIDKey ctxKey = "user_id"

// Auth middleware create.
func Auth(next http.Handler, jwtSecret []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.New()
		ctx := r.Context()

		header := r.Header.Get("Authorization")

		tokenStr := strings.TrimPrefix(header, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			return jwtSecret, nil
		})
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "cant parse token")
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		id, ok := claims["user_id"].(float64)
		if !ok {
			logger.Println("error with id in float64")
			return
		}

		ctx = context.WithValue(ctx, userIDKey, int64(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ContextFromUserID create.
func ContextFromUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// UserIDFromContext create.
func UserIDFromContext(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(userIDKey).(int64)
	return id, ok
}
