package sendcoin

import (
	"avito/internal/api/auth/middleware"
	mock_sendcoin "avito/internal/api/send_coin/mocks"
	sendcoinservice "avito/internal/service/send_coin_service"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func TestHandler_Send(t *testing.T) {
	var (
		errWithSend = errors.New("error with senddd")
	)
	type mockBehaviour func(sendService *mock_sendcoin.MockSendCoinService)
	logger, _ := test.NewNullLogger()

	senderID := int64(1)

	testCases := []struct {
		name      string
		senderID  int64
		inputBody string

		mockBehaviour mockBehaviour

		expectedCode int
	}{
		{
			name:      "OK",
			senderID:  senderID,
			inputBody: `{"toUser":"muhammad", "amount":100}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
				sendService.EXPECT().Send(gomock.Any(), senderID, "muhammad", int64(100))
			},
			expectedCode: http.StatusOK,
		},
		{
			name:      "send ERR",
			senderID:  senderID,
			inputBody: `{"toUser":"muhammad", "amount":100}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
				sendService.EXPECT().Send(gomock.Any(), senderID, "muhammad", int64(100)).Return(errWithSend)
			},
			expectedCode: http.StatusInternalServerError,
		},

		{
			name:      "bad request",
			senderID:  senderID,
			inputBody: `{"toUser":"muhammad", "amount":"asd"}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
			},

			expectedCode: http.StatusBadRequest,
		},
		{
			name:     "amount is negative",
			senderID: senderID,
			inputBody: `{
			"toUser":"muhammad",
			"amount":-1
			}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
				sendService.EXPECT().Send(gomock.Any(), senderID, "muhammad", int64(-1)).Return(sendcoinservice.ErrInvalidAmount)
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name:     "error with context",
			senderID: int64(0),
			inputBody: `{
			"toUser": "muhammad", 
			"amount":100
			}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
			},
			expectedCode: http.StatusInternalServerError,
		},

		{
			name:     "amount is missing",
			senderID: senderID,
			inputBody: `{
			"toUser":"muhammad"
			}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name:     "empty toUser",
			senderID: senderID,
			inputBody: `{
			"toUser":"",
			"amount":100
			}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
			},

			expectedCode: http.StatusBadRequest,
		},
		{
			name:     "amount is 0",
			senderID: senderID,
			inputBody: `{
			"toUser":"muhammad",
			"amount":0
			}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name:     "send to yourself",
			senderID: senderID,
			inputBody: `{
			"toUser":"muhammad",
			"amount":100
			}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
				sendService.EXPECT().
					Send(
						gomock.Any(),
						senderID,
						"muhammad",
						int64(100),
					).Return(sendcoinservice.ErrSelfTrans)
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:     "not enough coins",
			senderID: senderID,
			inputBody: `{
			"toUser":"muhammad",
			"amount":100
			}`,

			mockBehaviour: func(sendService *mock_sendcoin.MockSendCoinService) {
				sendService.EXPECT().
					Send(
						gomock.Any(),
						senderID,
						"muhammad",
						int64(100),
					).Return(sendcoinservice.ErrWithEnoughCoins)
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			sendService := mock_sendcoin.NewMockSendCoinService(ctrl)

			tc.mockBehaviour(sendService)

			handler := New(
				sendService,
				logger,
			)

			req := httptest.NewRequest(
				http.MethodPost,
				"/api/send",
				strings.NewReader(tc.inputBody),
			)

			var ctx context.Context

			if tc.name == "error with context" {
				ctx = req.Context()
			} else {
				ctx = middleware.ContextFromUserID(
					req.Context(),
					tc.senderID,
				)
			}

			req = req.WithContext(ctx)

			rec := httptest.NewRecorder()

			handler.Send(rec, req)

			require.Equal(
				t,
				tc.expectedCode,
				rec.Code,
			)
		})
	}
}
