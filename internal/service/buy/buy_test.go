package buyservice

import (
	"avito/internal/model"
	mock_buyservice "avito/internal/service/buy/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService_Buy(t *testing.T) {

	var (
		errGetMerchByName  = errors.New("error with get merch by name")
		errWithBuyMerch    = errors.New("error with buy merch")
		errWithGetUserByID = errors.New("error with get user by id")
	)

	userID := int64(1)

	testUser := model.User{
		ID:    userID,
		Coins: 500,
	}

	testMerch := model.Merch{
		ID:    1,
		Name:  "t-shirt",
		Price: 80,
	}

	type mockBehaviour func(
		purchasesRepo *mock_buyservice.MockPurchasesRepository,
		merchRepo *mock_buyservice.MockMerchRepository,
		userRepo *mock_buyservice.MockUserRepository,
	)

	testCases := []struct {
		name     string
		userID   int64
		itemName string

		mockBehaviour mockBehaviour

		expectedErr error
	}{
		{
			name:     "merch get OK",
			userID:   testMerch.ID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				purchasesRepo *mock_buyservice.MockPurchasesRepository,
				merchRepo *mock_buyservice.MockMerchRepository,
				userRepo *mock_buyservice.MockUserRepository,
			) {
				merchRepo.EXPECT().GetByName(gomock.Any(), testMerch.Name).Return(testMerch, nil)
				purchasesRepo.EXPECT().Buy(gomock.Any(), userID, testMerch.ID, testMerch.Price).Return(nil)
				userRepo.EXPECT().GetByID(gomock.Any(), testUser.ID).Return(testUser, nil)
			},

			expectedErr: nil,
		},
		{
			name:     "buy OK",
			userID:   userID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				purchasesRepo *mock_buyservice.MockPurchasesRepository,
				merchRepo *mock_buyservice.MockMerchRepository,
				userRepo *mock_buyservice.MockUserRepository,
			) {
				merchRepo.EXPECT().GetByName(gomock.Any(), testMerch.Name).Return(testMerch, nil)
				purchasesRepo.EXPECT().Buy(gomock.Any(), userID, testMerch.ID, testMerch.Price).Return(nil)
				userRepo.EXPECT().GetByID(gomock.Any(), testUser.ID).Return(testUser, nil)
			},

			expectedErr: nil,
		},
		{
			name:     "merch get ERR",
			userID:   userID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				purchasesRepo *mock_buyservice.MockPurchasesRepository,
				merchRepo *mock_buyservice.MockMerchRepository,
				userRepo *mock_buyservice.MockUserRepository,
			) {
				merchRepo.EXPECT().GetByName(gomock.Any(), testMerch.Name).Return(model.Merch{}, errGetMerchByName)
			},
			expectedErr: errGetMerchByName,
		},
		{
			name:     "empty merch's name",
			userID:   userID,
			itemName: "",

			mockBehaviour: func(
				purchasesRepo *mock_buyservice.MockPurchasesRepository,
				merchRepo *mock_buyservice.MockMerchRepository,
				userRepo *mock_buyservice.MockUserRepository,
			) {
			},

			expectedErr: ErrWithItemName,
		},
		{
			name:     "buy ERR",
			userID:   userID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				purchasesRepo *mock_buyservice.MockPurchasesRepository,
				merchRepo *mock_buyservice.MockMerchRepository,
				userRepo *mock_buyservice.MockUserRepository,
			) {
				merchRepo.EXPECT().GetByName(gomock.Any(), testMerch.Name).Return(testMerch, nil)
				userRepo.EXPECT().GetByID(gomock.Any(), testUser.ID).Return(testUser, nil)
				purchasesRepo.EXPECT().Buy(gomock.Any(), testUser.ID, testMerch.ID, testMerch.Price).Return(errWithBuyMerch)
			},
			expectedErr: errWithBuyMerch,
		},
		{
			name:     "not enough coins",
			userID:   userID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				purchasesRepo *mock_buyservice.MockPurchasesRepository,
				merchRepo *mock_buyservice.MockMerchRepository,
				userRepo *mock_buyservice.MockUserRepository) {

				poorUser := testUser
				poorUser.Coins = 50

				merchRepo.EXPECT().GetByName(gomock.Any(), testMerch.Name).Return(testMerch, nil)
				userRepo.EXPECT().GetByID(gomock.Any(), poorUser.ID).Return(poorUser, nil)
				purchasesRepo.EXPECT().Buy(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			expectedErr: ErrWithEnoughCoins,
		},
		{
			name:     "user get ERR",
			userID:   userID,
			itemName: testMerch.Name,

			mockBehaviour: func(purchasesRepo *mock_buyservice.MockPurchasesRepository,
				merchRepo *mock_buyservice.MockMerchRepository,
				userRepo *mock_buyservice.MockUserRepository,
			) {
				merchRepo.EXPECT().GetByName(gomock.Any(), testMerch.Name).Return(testMerch, nil)
				userRepo.EXPECT().GetByID(gomock.Any(), testUser.ID).Return(model.User{}, errWithGetUserByID)
			},

			expectedErr: errWithGetUserByID,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			purchasesRepo := mock_buyservice.NewMockPurchasesRepository(ctrl)
			merchRepo := mock_buyservice.NewMockMerchRepository(ctrl)
			userRepo := mock_buyservice.NewMockUserRepository(ctrl)

			tc.mockBehaviour(purchasesRepo, merchRepo, userRepo)

			service := New(
				purchasesRepo,
				merchRepo,
				userRepo,
			)

			err := service.Buy(context.Background(), tc.userID, tc.itemName)

			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}

		})
	}
}
