package service_test

import (
	"avito/internal/repository/merch"
	"avito/internal/repository/purchases"
	"avito/internal/repository/users"
	"avito/internal/service"
	"avito/internal/service/mocks"
	"context"
	"testing"

	"github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestService_Buy_Integration_Test create.
func TestService_Buy_Integration_Test(t *testing.T) {
	// Merch repo, purch repo i user repo
	pool := newTestDB(t)
	logger := logrus.New()
	ctx := context.Background()
	transactor := manager.Must(
		pgxv5.NewDefaultFactory(pool),
	)

	purchRepo := purchases.New(
		pgxv5.DefaultCtxGetter,
		pool,
		logger,
	)

	merchRepo := merch.New(
		pool,
		logger,
	)

	userRepo := users.New(
		pool,
		logger,
		pgxv5.DefaultCtxGetter,
	)

	ctrl := gomock.NewController(t)

	produce := mocks.NewMockProduce(ctrl)

	produce.EXPECT().Produce(
		gomock.Any(),
		gomock.Any(),
	).Return(nil)

	service := service.New(
		&userRepo,
		nil,
		&purchRepo,
		&merchRepo,
		transactor,
		produce,
		"",
	)

	user, err := userRepo.Create(
		ctx,
		"muhammad",
		"muhammad@mail.ru",
		"hash",
	)
	require.NoError(t, err)

	err = service.Buy(
		ctx,
		user.ID,
		"book",
	)
	require.NoError(t, err)

	userAfter, err := userRepo.GetByID(
		ctx,
		user.ID,
	)
	require.NoError(t, err)

	assert.Equal(
		t,
		int64(950),
		userAfter.Coins,
	)
}
