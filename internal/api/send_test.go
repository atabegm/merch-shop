package api

import (
	"avito/internal/api/middleware"
	mock_api "avito/internal/api/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
)

func TestHandler_Send(t *testing.T) {
	var (
		errWithSend = errors.New("error with senddd")
	)
	type mockBehaviour func(service *mock_api.MockService)
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

			mockBehaviour: func(service *mock_api.MockService) {
				service.EXPECT().SendCoins(
					gomock.Any(),
					senderID,
					"muhammad",
					"100",
				).Return(nil)
			},
			expectedCode: http.StatusOK,
		},

		{
			name:     "send ERR",
			senderID: senderID,
			inputBody: `{
				"toUser":"muhammad",
				"amount":100
			}`,

			mockBehaviour: func(service *mock_api.MockService) {
				service.EXPECT().SendCoins(
					gomock.Any(),
					senderID,
					"muhammad",
					"100",
				).Return(errWithSend)
			},
			expectedCode: http.StatusInternalServerError,
		},

		{
			name:      "bad request",
			senderID:  senderID,
			inputBody: `{"toUser":"muhammad", "amount":"asd"}`,

			mockBehaviour: func(service *mock_api.MockService) {

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

			mockBehaviour: func(service *mock_api.MockService) {
				service.EXPECT().SendCoins(
					gomock.Any(),
					senderID,
					"muhammad",
					"100",
				).Return(nil)
			},
			expectedCode: http.StatusInternalServerError,
		},

		{
			name:     "send to yourself",
			senderID: senderID,
			inputBody: `{
			"toUser":"muhammad",
			"amount":100
			}`,

			mockBehaviour: func(service *mock_api.MockService) {
				service.EXPECT().SendCoins(
					gomock.Any(),
					senderID,
					"muhammad",
					"100",
				).Return(nil)
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

			mockBehaviour: func(service *mock_api.MockService) {
				service.EXPECT().SendCoins(
					gomock.Any(),
					senderID,
					"muhammad",
					"100",
				).Return(nil)
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			service := mock_api.NewMockService(ctrl)

			tc.mockBehaviour(service)

			handler := New(
				service,
				logger,
			)

			req := httptest.NewRequest(
				http.MethodPost,
				"/api/send",
				strings.NewReader(tc.inputBody),
			)

			ctx := middleware.ContextFromUserID(
				req.Context(),
				tc.senderID,
			)

			req = req.WithContext(ctx)

			rec := httptest.NewRecorder()

			handler.Send(
				rec,
				req,
			)

			
		})
	}
}
