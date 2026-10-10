package service_test

import (
	"avito/internal/repository/users"
	"avito/internal/service"
	"context"
	"testing"

	"github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const databaseURL = "postgres://postgres:password@localhost:5433/shop?sslmode=disable"

// newTestDB create.
func newTestDB(t *testing.T) *pgxpool.Pool {
	// t.Helper()
	// ctx := context.Background()

	// pool, err := pgxpool.New(
	// 	ctx,
	// 	databaseURL,
	// )

	// require.NoError(
	// 	t,
	// 	err,
	// )

	// _, err = pool.Exec(
	// 	ctx,
	// 	"TRUNCATE TABLE users RESTART IDENTITY CASCADE",
	// )
	// assert.NoError(
	// 	t,
	// 	err,
	// )

	// t.Cleanup(func() {
	// 	pool.Close()
	// })

	// return pool

	t.Helper()

	ctx := context.Background()

	pool, err := pgxpool.New(
		ctx,
		databaseURL,
	)
	require.NoError(
		t,
		err,
	)

	_, err = pool.Exec(
		ctx,
		"TRUNCATE TABLE users RESTART IDENTITY CASCADE",
	)
	require.NoError(
		t,
		err,
	)

	t.Cleanup(func() {
		pool.Close()
	})

	return pool
}

// TestService_SendCoins_Integration create.
func TestService_SendCoins_Integration(t *testing.T) {
	pool := newTestDB(t)

	ctx := context.Background()

	logger := logrus.New()

	userRepo := users.New(
		pool,
		logger,
		pgxv5.DefaultCtxGetter,
	)

	transactor := manager.Must(
		pgxv5.NewDefaultFactory(pool),
	)

	service := service.New(
		&userRepo,
		nil,
		nil,
		nil,
		transactor,
		nil,
		"",
	)

	sender, err := userRepo.Create(
		ctx,
		"muhammad",
		"muhammad@mail.ru",
		"hash-password",
	)
	assert.NoError(t, err)

	receiver, err := userRepo.Create(
		ctx,
		"ibrahim",
		"ibrahim@mail.ru",
		"hash-password",
	)
	assert.NoError(t, err)

	err = service.SendCoins(
		ctx,
		sender.ID,
		receiver.Username,
		int64(100),
	)
	assert.NoError(t, err)

	sender2, err := userRepo.GetByID(
		ctx,
		sender.ID,
	)
	assert.NoError(t, err)

	receiver2, err := userRepo.GetByID(
		ctx,
		receiver.ID,
	)
	assert.NoError(t, err)

	assert.Equal(
		t,
		int64(900),
		sender2.Coins,
	)

	assert.Equal(
		t,
		int64(1100),
		receiver2.Coins,
	)
}
