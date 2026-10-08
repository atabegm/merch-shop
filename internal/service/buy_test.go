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
	errMerch       = errors.New("merch error")
	errUser        = errors.New("user error")
	errPurchase    = errors.New("purchase error")
	errCoins       = errors.New("coins error")
	errTransaction = errors.New("transaction error")
	errProduce     = errors.New("produce error")
)

func TestService_Buy(t *testing.T) {
	type mockBehaviour func(
		merchRepo *mock_service.MockMerchRepository,
		userRepo *mock_service.MockUserRepository,
		purchRepo *mock_service.MockPurchasesRepository,
		transactor *mock_service.MockTransactor,
		produce *mock_service.MockProduce,
	)

	testMerch := model.Merch{
		ID:    1,
		Name:  "book",
		Price: 100,
	}

	testUser := model.User{
		ID:    1,
		Email: "muhammad@mail.ru",
	}

	testCases := []struct {
		name     string
		userID   int64
		itemName string

		mockBehaviour mockBehaviour

		expectedErr error
	}{
		{
			name:     "buy OK",
			userID:   testUser.ID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				merchRepo *mock_service.MockMerchRepository,
				userRepo *mock_service.MockUserRepository,
				purchRepo *mock_service.MockPurchasesRepository,
				transactor *mock_service.MockTransactor,
				produce *mock_service.MockProduce,
			) {
				merchRepo.EXPECT().
					GetByName(
						gomock.Any(),
						testMerch.Name,
					).
					Return(testMerch, nil)

				userRepo.EXPECT().
					GetByID(
						gomock.Any(),
						testUser.ID,
					).
					Return(testUser, nil)

				transactor.EXPECT().
					Do(
						gomock.Any(),
						gomock.Any(),
					).
					DoAndReturn(
						func(
							ctx context.Context,
							fn func(context.Context) error,
						) error {
							return fn(ctx)
						},
					)

				purchRepo.EXPECT().
					Create(
						gomock.Any(),
						testUser.ID,
						testMerch.ID,
					).
					Return(nil)

				userRepo.EXPECT().
					SubstractCoins(
						gomock.Any(),
						testMerch.Price,
						testUser.ID,
					).
					Return(nil)

				produce.EXPECT().
					Produce(
						gomock.Any(),
						gomock.Any(),
					).
					Return(nil)
			},

			expectedErr: nil,
		},

		{
			name:     "empty item name",
			userID:   testUser.ID,
			itemName: "",

			mockBehaviour: func(
				merchRepo *mock_service.MockMerchRepository,
				userRepo *mock_service.MockUserRepository,
				purchRepo *mock_service.MockPurchasesRepository,
				transactor *mock_service.MockTransactor,
				produce *mock_service.MockProduce,
			) {
			},

			expectedErr: ErrEmptyItem,
		},

		{
			name:     "get merch error",
			userID:   testUser.ID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				merchRepo *mock_service.MockMerchRepository,
				userRepo *mock_service.MockUserRepository,
				purchRepo *mock_service.MockPurchasesRepository,
				transactor *mock_service.MockTransactor,
				produce *mock_service.MockProduce,
			) {
				merchRepo.EXPECT().
					GetByName(
						gomock.Any(),
						testMerch.Name,
					).
					Return(model.Merch{}, errMerch)
			},

			expectedErr: errMerch,
		},

		{
			name:     "get user error",
			userID:   testUser.ID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				merchRepo *mock_service.MockMerchRepository,
				userRepo *mock_service.MockUserRepository,
				purchRepo *mock_service.MockPurchasesRepository,
				transactor *mock_service.MockTransactor,
				produce *mock_service.MockProduce,
			) {
				merchRepo.EXPECT().
					GetByName(
						gomock.Any(),
						testMerch.Name,
					).
					Return(testMerch, nil)

				userRepo.EXPECT().
					GetByID(
						gomock.Any(),
						testUser.ID,
					).
					Return(model.User{}, errUser)
			},

			expectedErr: errUser,
		},

		{
			name:     "create purchase error",
			userID:   testUser.ID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				merchRepo *mock_service.MockMerchRepository,
				userRepo *mock_service.MockUserRepository,
				purchRepo *mock_service.MockPurchasesRepository,
				transactor *mock_service.MockTransactor,
				produce *mock_service.MockProduce,
			) {
				merchRepo.EXPECT().
					GetByName(gomock.Any(), testMerch.Name).
					Return(testMerch, nil)

				userRepo.EXPECT().
					GetByID(gomock.Any(), testUser.ID).
					Return(testUser, nil)

				transactor.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(
							ctx context.Context,
							fn func(context.Context) error,
						) error {
							return fn(ctx)
						},
					)

				purchRepo.EXPECT().
					Create(
						gomock.Any(),
						testUser.ID,
						testMerch.ID,
					).
					Return(errPurchase)
			},

			expectedErr: errPurchase,
		},

		{
			name:     "subtract coins error",
			userID:   testUser.ID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				merchRepo *mock_service.MockMerchRepository,
				userRepo *mock_service.MockUserRepository,
				purchRepo *mock_service.MockPurchasesRepository,
				transactor *mock_service.MockTransactor,
				produce *mock_service.MockProduce,
			) {
				merchRepo.EXPECT().
					GetByName(gomock.Any(), testMerch.Name).
					Return(testMerch, nil)

				userRepo.EXPECT().
					GetByID(gomock.Any(), testUser.ID).
					Return(testUser, nil)

				transactor.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(
							ctx context.Context,
							fn func(context.Context) error,
						) error {
							return fn(ctx)
						},
					)

				purchRepo.EXPECT().
					Create(
						gomock.Any(),
						testUser.ID,
						testMerch.ID,
					).
					Return(nil)

				userRepo.EXPECT().
					SubstractCoins(
						gomock.Any(),
						testMerch.Price,
						testUser.ID,
					).
					Return(errCoins)
			},

			expectedErr: errCoins,
		},

		{
			name:     "transaction error",
			userID:   testUser.ID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				merchRepo *mock_service.MockMerchRepository,
				userRepo *mock_service.MockUserRepository,
				purchRepo *mock_service.MockPurchasesRepository,
				transactor *mock_service.MockTransactor,
				produce *mock_service.MockProduce,
			) {
				merchRepo.EXPECT().
					GetByName(gomock.Any(), testMerch.Name).
					Return(testMerch, nil)

				userRepo.EXPECT().
					GetByID(gomock.Any(), testUser.ID).
					Return(testUser, nil)

				transactor.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					Return(errTransaction)
			},

			expectedErr: errTransaction,
		},
		{
			name:     "produce error",
			userID:   testUser.ID,
			itemName: testMerch.Name,

			mockBehaviour: func(
				merchRepo *mock_service.MockMerchRepository,
				userRepo *mock_service.MockUserRepository,
				purchRepo *mock_service.MockPurchasesRepository,
				transactor *mock_service.MockTransactor,
				produce *mock_service.MockProduce,
			) {
				merchRepo.EXPECT().
					GetByName(gomock.Any(), testMerch.Name).
					Return(testMerch, nil)

				userRepo.EXPECT().
					GetByID(gomock.Any(), testUser.ID).
					Return(testUser, nil)

				transactor.EXPECT().
					Do(gomock.Any(), gomock.Any()).
					DoAndReturn(
						func(
							ctx context.Context,
							fn func(context.Context) error,
						) error {
							return fn(ctx)
						},
					)

				purchRepo.EXPECT().
					Create(
						gomock.Any(),
						testUser.ID,
						testMerch.ID,
					).
					Return(nil)

				userRepo.EXPECT().
					SubstractCoins(
						gomock.Any(),
						testMerch.Price,
						testUser.ID,
					).
					Return(nil)

				produce.EXPECT().
					Produce(
						gomock.Any(),
						gomock.Any(),
					).
					Return(errProduce)
			},

			expectedErr: errProduce,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			userRepo := mock_service.NewMockUserRepository(ctrl)
			merchRepo := mock_service.NewMockMerchRepository(ctrl)
			purchRepo := mock_service.NewMockPurchasesRepository(ctrl)
			transactor := mock_service.NewMockTransactor(ctrl)
			produce := mock_service.NewMockProduce(ctrl)

			tc.mockBehaviour(
				merchRepo,
				userRepo,
				purchRepo,
				transactor,
				produce,
			)

			service := &Service{
				UserRepository:      userRepo,
				MerchRepository:     merchRepo,
				PurchasesRepository: purchRepo,
				Transactor:          transactor,
				Produce:             produce,
			}

			err := service.Buy(
				context.Background(),
				tc.userID,
				tc.itemName,
			)

			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
