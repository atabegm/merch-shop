package service

import (
	"avito/internal/model"
	mock_service "avito/internal/service/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	errGetCoins     = errors.New("get coins error")
	errGetInventory = errors.New("get inventory error")
	errGetReceived  = errors.New("get received error")
	errGetSent      = errors.New("get sent error")
)

func TestService_Info(t *testing.T) {
	userID := int64(1)

	type mockBehavior func(
		userRepo *mock_service.MockUserRepository,
		purchasesRepo *mock_service.MockPurchasesRepository,
		coinTransfersRepo *mock_service.MockCoinTransfersRepository,
	)

	testCases := []struct {
		name string

		mockBehavior mockBehavior

		expectedErr error
	}{
		{
			name: "info OK",

			mockBehavior: func(
				userRepo *mock_service.MockUserRepository,
				purchasesRepo *mock_service.MockPurchasesRepository,
				coinTransfersRepo *mock_service.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().
					GetCoins(
						gomock.Any(),
						userID,
					).
					Return(int64(1000), nil)

				purchasesRepo.EXPECT().
					GetInventory(
						gomock.Any(),
						userID,
					).
					Return(nil, nil)

				coinTransfersRepo.EXPECT().
					GetReceived(
						gomock.Any(),
						userID,
					).
					Return(nil, nil)

				coinTransfersRepo.EXPECT().
					GetSent(
						gomock.Any(),
						userID,
					).
					Return(nil, nil)
			},
		},

		{
			name: "get coins error",

			mockBehavior: func(
				userRepo *mock_service.MockUserRepository,
				purchasesRepo *mock_service.MockPurchasesRepository,
				coinTransfersRepo *mock_service.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().
					GetCoins(
						gomock.Any(),
						userID,
					).
					Return(int64(0), errGetCoins)
			},

			expectedErr: errGetCoins,
		},

		{
			name: "get inventory error",

			mockBehavior: func(
				userRepo *mock_service.MockUserRepository,
				purchasesRepo *mock_service.MockPurchasesRepository,
				coinTransfersRepo *mock_service.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().
					GetCoins(
						gomock.Any(),
						userID,
					).
					Return(int64(1000), nil)

				purchasesRepo.EXPECT().
					GetInventory(
						gomock.Any(),
						userID,
					).
					Return(nil, errGetInventory)
			},

			expectedErr: errGetInventory,
		},

		{
			name: "get received error",

			mockBehavior: func(
				userRepo *mock_service.MockUserRepository,
				purchasesRepo *mock_service.MockPurchasesRepository,
				coinTransfersRepo *mock_service.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().
					GetCoins(
						gomock.Any(),
						userID,
					).
					Return(int64(1000), nil)

				purchasesRepo.EXPECT().
					GetInventory(
						gomock.Any(),
						userID,
					).
					Return(nil, nil)

				coinTransfersRepo.EXPECT().
					GetReceived(
						gomock.Any(),
						userID,
					).
					Return(nil, errGetReceived)
			},

			expectedErr: errGetReceived,
		},

		{
			name: "get sent error",

			mockBehavior: func(
				userRepo *mock_service.MockUserRepository,
				purchasesRepo *mock_service.MockPurchasesRepository,
				coinTransfersRepo *mock_service.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().
					GetCoins(
						gomock.Any(),
						userID,
					).
					Return(int64(1000), nil)

				purchasesRepo.EXPECT().
					GetInventory(
						gomock.Any(),
						userID,
					).
					Return(nil, nil)

				coinTransfersRepo.EXPECT().
					GetReceived(
						gomock.Any(),
						userID,
					).
					Return(nil, nil)

				coinTransfersRepo.EXPECT().
					GetSent(
						gomock.Any(),
						userID,
					).
					Return(nil, errGetSent)
			},

			expectedErr: errGetSent,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			userRepo := mock_service.NewMockUserRepository(ctrl)
			purchasesRepo := mock_service.NewMockPurchasesRepository(ctrl)
			coinTransfersRepo := mock_service.NewMockCoinTransfersRepository(ctrl)

			tc.mockBehavior(
				userRepo,
				purchasesRepo,
				coinTransfersRepo,
			)

			service := &Service{
				UserRepository:          userRepo,
				PurchasesRepository:     purchasesRepo,
				CoinTransfersRepository: coinTransfersRepo,
			}

			info, err := service.Info(
				context.Background(),
				userID,
			)

			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
				require.Equal(t, model.Info{}, info)
				return
			}

			require.NoError(t, err)
			require.Equal(t, int64(1000), info.Coins)
		})
	}
}
