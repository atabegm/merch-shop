package sendcoinservice

import (
	"avito/internal/model"
	mock_sendcoinservice "avito/internal/service/send_coin_service/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// TestService_SendCoin create.
func TestService_SendCoin(t *testing.T) {
	var (
		errWithGetSenderByID         = errors.New("error with get sender by id")
		errWithGetReceiverByUsername = errors.New("error with get receiver by username")
		errWithSend                  = errors.New("error with send")
	)

	type mockBehaviour func(
		userRepo *mock_sendcoinservice.MockUserRepository,
		coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
	)

	senderID := int64(1)
	receiverID := int64(2)

	amount := int64(100)

	receiverUser := model.User{
		ID:       receiverID,
		Username: "muhammad",
	}

	senderUser := model.User{
		ID: senderID,
	}

	testCases := []struct {
		name     string
		senderID int64
		toUser   string
		amount   int64

		mockBehaviour mockBehaviour

		expectedErr error
	}{
		{
			name:     "send OK",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(receiverUser, nil)
				userRepo.EXPECT().GetByID(gomock.Any(), senderID).Return(senderUser, nil)
				coinTransfersRepo.EXPECT().Send(gomock.Any(), senderUser.ID, receiverUser.ID, amount).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:     "send ERR",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(receiverUser, nil)
				userRepo.EXPECT().GetByID(gomock.Any(), senderID).Return(senderUser, nil)
				coinTransfersRepo.EXPECT().Send(gomock.Any(), senderUser.ID, receiverUser.ID, amount).Return(errWithSend)
			},
			expectedErr: errWithSend,
		},
		{
			name:     "negative amount",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   int64(-10),

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
			},
			expectedErr: ErrInvalidAmount,
		},

		{
			name:     "get sender by id ERR",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().GetByID(gomock.Any(), senderID).Return(model.User{}, errWithGetSenderByID)
			},

			expectedErr: errWithGetSenderByID,
		},
		{
			name:     "get sender by id OK",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().GetByID(gomock.Any(), senderID).Return(senderUser, nil)
				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(receiverUser, nil)
				coinTransfersRepo.EXPECT().Send(gomock.Any(), senderID, receiverUser.ID, amount).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:     "get receiver by id OK",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().GetByID(gomock.Any(), senderID).Return(senderUser, nil)
				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(receiverUser, nil)
				coinTransfersRepo.EXPECT().Send(gomock.Any(), senderID, receiverUser.ID, amount).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:     "get receiver by id ERR",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().GetByID(gomock.Any(), senderID).Return(senderUser, nil)
				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(receiverUser, nil)
				coinTransfersRepo.EXPECT().Send(gomock.Any(), senderID, receiverUser.ID, amount).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:     "get receiver by username OK",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().GetByID(gomock.Any(), senderID).Return(senderUser, nil)
				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(receiverUser, nil)
				coinTransfersRepo.EXPECT().Send(gomock.Any(), senderID, receiverUser.ID, amount).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name:     "get receiver by username ERR",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().GetByID(gomock.Any(), senderID).Return(senderUser, nil)
				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(model.User{}, errWithGetReceiverByUsername)
			},
			expectedErr: errWithGetReceiverByUsername,
		},
		{
			name:     "self to urself",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_sendcoinservice.MockUserRepository,
				coinTransfersRepo *mock_sendcoinservice.MockCoinTransfersRepository,
			) {
				userRepo.EXPECT().GetByID(gomock.Any(), senderID).Return(receiverUser, nil)
				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(receiverUser, nil)
				coinTransfersRepo.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			expectedErr: ErrSelfTrans,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			userRepo := mock_sendcoinservice.NewMockUserRepository(ctrl)
			coinTransfersRepo := mock_sendcoinservice.NewMockCoinTransfersRepository(ctrl)

			tc.mockBehaviour(
				userRepo,
				coinTransfersRepo,
			)

			service := New(userRepo, coinTransfersRepo)

			err := service.Send(context.Background(), tc.senderID, tc.toUser, tc.amount)
			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
