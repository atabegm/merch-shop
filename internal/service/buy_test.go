package service

// import (
// 	mock_service "avito/internal/service/mocks"
// 	"context"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// func TestService_Buy(t *testing.T) {
// 	type mockBehaviour func(
// 		userRepo *mock_service.MockUserRepository,
// 		purchasesRepo *mock_service.MockPurchasesRepository,
// 		merchRepo *mo
// 		transactor *mock_service.MockTransactor,
// 	)

// 	userID := int64(1)

// 	itemName := "book"
// 	quantity := int64(1)

// 	amount := int64(1)

// 	testCases := []struct {
// 		name     string
// 		userID   int64
// 		itemName string
// 		quantity int64
// 		amount   int64

// 		mockBehaviour mockBehaviour

// 		expectedErr error
// 	}{
// 		{
// 			name:     "buy OK",
// 			userID:   userID,
// 			itemName: itemName,
// 			quantity: quantity,
// 			amount:   amount,

// 			mockBehaviour: func(
// 				userRepo *mock_service.MockUserRepository,
// 				purchasesRepo *mock_service.MockPurchasesRepository,
// 				transactor *mock_service.MockTransactor,
// 			) {
// 				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
// 					func(
// 						ctx context.Context,
// 						fn func(context.Context) error,
// 					) error {
// 						return fn(ctx)
// 					},
// 				)

// 				userRepo.EXPECT().SubstractCoins(
// 					gomock.Any(),
// 					amount,
// 					userID,
// 				)

// 				purchasesRepo.EXPECT().Create(
// 					context.Background(),
// 					userID,

// 				)
// 			},

// 			expectedErr: nil,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)

// 			userRepo := mock_service.NewMockUserRepository(ctrl)
// 			purchasesRepo := mock_service.NewMockPurchasesRepository(ctrl)
// 			transactor := mock_service.NewMockTransactor(ctrl)

// 			tc.mockBehaviour(
// 				userRepo,
// 				purchasesRepo,
// 				transactor,
// 			)

// 			buyService := &Service{
// 				UserRepository:      userRepo,
// 				PurchasesRepository: purchasesRepo,
// 				Transactor:          transactor,
// 			}

// 			err := buyService.Buy(
// 				context.Background(),
// 				tc.userID,
// 				tc.itemName,
// 				tc.quantity,
// 			)

// 			if tc.expectedErr != nil {
// 				require.ErrorIs(
// 					t,
// 					err,
// 					tc.expectedErr,
// 				)
// 			} else {
// 				require.NoError(
// 					t,
// 					err,
// 				)
// 			}
// 		})
// 	}
// }
