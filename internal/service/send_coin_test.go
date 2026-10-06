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

// TestService_SendCoin create.
func TestService_SendCoin(t *testing.T) {
	var (
		errSubCoins                  = errors.New("error with sub coins")
		errWithGetReceiverByUsername = errors.New("error with get receiver by username")
		errWithTransactor            = errors.New("error with transactor")
		errAddCoins                  = errors.New("error with add coins")
	)

	type mockBehaviour func(
		userRepo *mock_service.MockUserRepository,
		transactor *mock_service.MockTransactor,
	)

	senderID := int64(1)
	receiverID := int64(2)

	amount := int64(100)

	receiverUser := model.User{
		ID:       receiverID,
		Username: "muhammad",
	}

	sendUser := model.User{
		ID:       senderID,
		Username: "ibrahim",
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
				userRepo *mock_service.MockUserRepository,
				transactor *mock_service.MockTransactor,
			) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(
						ctx context.Context,
						fn func(context.Context) error,
					) error {
						return fn(ctx)
					},
				)
				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(receiverUser, nil)
				userRepo.EXPECT().AddCoins(
					gomock.Any(),
					amount,
					receiverUser.ID,
				).Return(nil)

				userRepo.EXPECT().SubstractCoins(
					gomock.Any(),
					amount,
					senderID,
				)
			},

			expectedErr: nil,
		},
		{
			name:     "send ERR",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_service.MockUserRepository,
				transactor *mock_service.MockTransactor,
			) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).Return(errWithTransactor)
			},
			expectedErr: errWithTransactor,
		},
		{
			name:     "negative amount",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   int64(-10),

			mockBehaviour: func(
				userRepo *mock_service.MockUserRepository,
				transactor *mock_service.MockTransactor,
			) {
			},
			expectedErr: ErrEmptyAmount,
		},

		{
			name:     "get by username ERR",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_service.MockUserRepository,
				transactor *mock_service.MockTransactor,
			) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(
						ctx context.Context,
						fn func(context.Context) error,
					) error {
						return fn(ctx)
					},
				)

				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(model.User{}, errWithGetReceiverByUsername)
			},

			expectedErr: errWithGetReceiverByUsername,
		},
		{
			name:     "get by username OK",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_service.MockUserRepository,
				transactor *mock_service.MockTransactor,
			) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(
						ctx context.Context,
						fn func(context.Context) error,
					) error {
						return fn(ctx)
					},
				)

				userRepo.EXPECT().GetByUsername(gomock.Any(), receiverUser.Username).Return(
					receiverUser,
					nil,
				)

				userRepo.EXPECT().AddCoins(
					gomock.Any(),
					amount,
					receiverUser.ID,
				).Return(nil)

				userRepo.EXPECT().SubstractCoins(
					gomock.Any(),
					amount,
					senderID,
				).Return(nil)
			},

			expectedErr: nil,
		},

		{
			name:     "self to urself",
			senderID: senderID,
			toUser:   sendUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_service.MockUserRepository,
				transactor *mock_service.MockTransactor,
			) {
				transactor.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(
					func(
						ctx context.Context,
						fn func(context.Context) error,
					) error {
						return fn(ctx)
					},
				)

				userRepo.EXPECT().GetByUsername(
					gomock.Any(),
					sendUser.Username,
				).Return(sendUser, nil)
			},
			expectedErr: ErrSelfTrans,
		},
		{
			name:     "add coins ERR",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_service.MockUserRepository,
				transactor *mock_service.MockTransactor,
			) {
				transactor.EXPECT().Do(
					gomock.Any(),
					gomock.Any(),
				).DoAndReturn(
					func(
						ctx context.Context,
						fn func(context.Context) error,
					) error {
						return fn(ctx)
					},
				)

				userRepo.EXPECT().GetByUsername(
					gomock.Any(),
					receiverUser.Username,
				).Return(receiverUser, nil)

				userRepo.EXPECT().AddCoins(
					gomock.Any(),
					amount,
					receiverUser.ID,
				).Return(errAddCoins)
			},

			expectedErr: errAddCoins,
		},

		{
			name:     "subtract coins ERR",
			senderID: senderID,
			toUser:   receiverUser.Username,
			amount:   amount,

			mockBehaviour: func(
				userRepo *mock_service.MockUserRepository,
				transactor *mock_service.MockTransactor,
			) {
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

				userRepo.EXPECT().
					GetByUsername(
						gomock.Any(),
						receiverUser.Username,
					).
					Return(receiverUser, nil)

				userRepo.EXPECT().
					AddCoins(
						gomock.Any(),
						amount,
						receiverUser.ID,
					).
					Return(nil)

				userRepo.EXPECT().
					SubstractCoins(
						gomock.Any(),
						amount,
						senderID,
					).
					Return(errSubCoins)
			},

			expectedErr: errSubCoins,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			userRepo := mock_service.NewMockUserRepository(ctrl)
			transactor := mock_service.NewMockTransactor(ctrl)

			tc.mockBehaviour(
				userRepo,
				transactor,
			)

			service := &Service{
				UserRepository: userRepo,
				Transactor:     transactor,
			}

			err := service.SendCoins(
				context.Background(),
				tc.senderID,
				tc.toUser,
				tc.amount,
			)

			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
