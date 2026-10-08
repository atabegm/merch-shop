package api

import (
	"avito/internal/api/middleware"
	mock_api "avito/internal/api/mocks"
	"avito/internal/service"
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
	logger, _ := test.NewNullLogger()
	senderID := int64(1)

	errService := errors.New("service error")

	type mockBehavior func(
		serviceMock *mock_api.MockService,
	)

	testCases := []struct {
		name string
		body string

		contextFromUserID bool

		mockBehavior mockBehavior

		expectedStatus int
	}{
		{
			name: "send OK",
			body: `{
				"toUser":"muhammad",
				"amount":100
			}`,

			contextFromUserID: true,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().
					SendCoins(
						gomock.Any(),
						senderID,
						"muhammad",
						int64(100),
					).Return(nil)
			},

			expectedStatus: http.StatusOK,
		},

		{
			name: "invalid json",
			body: `{`,

			contextFromUserID: true,

			mockBehavior: func(serviceMock *mock_api.MockService) {
			},

			expectedStatus: http.StatusBadRequest,
		},

		{
			name: "empty to user",
			body: `{
				"toUser":"",
				"amount":100
			}`,

			contextFromUserID: true,

			mockBehavior: func(serviceMock *mock_api.MockService) {
			},

			expectedStatus: http.StatusBadRequest,
		},

		{
			name: "empty amount",
			body: `{
				"toUser":"muhammad",
				"amount":0
			}`,

			contextFromUserID: true,

			mockBehavior: func(serviceMock *mock_api.MockService) {
			},

			expectedStatus: http.StatusBadRequest,
		},

		{
			name: "self transfer",
			body: `{
				"toUser":"muhammad",
				"amount":100
			}`,

			contextFromUserID: true,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().
					SendCoins(
						gomock.Any(),
						senderID,
						"muhammad",
						int64(100),
					).Return(service.ErrSelfTrans)
			},

			expectedStatus: http.StatusBadRequest,
		},

		{
			name: "not enough coins",
			body: `{
				"toUser":"muhammad",
				"amount":100
			}`,

			contextFromUserID: true,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().
					SendCoins(
						gomock.Any(),
						senderID,
						"muhammad",
						int64(100),
					).Return(service.ErrNotEnoughCoins)
			},

			expectedStatus: http.StatusBadRequest,
		},

		{
			name: "send Error",
			body: `{
				"toUser":"muhammad",
				"amount":100
			}`,

			contextFromUserID: true,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().
					SendCoins(
						gomock.Any(),
						senderID,
						"muhammad",
						int64(100),
					).Return(errService)
			},

			expectedStatus: http.StatusInternalServerError,
		},

		{
			name: "sender ID not in context",
			body: `{
				"toUser":"muhammad",
				"amount":100
			}`,

			contextFromUserID: false,

			mockBehavior: func(serviceMock *mock_api.MockService) {
			},

			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			serviceMock := mock_api.NewMockService(ctrl)

			tc.mockBehavior(serviceMock)

			handler := &Handler{
				service: serviceMock,
				logger:  logger,
			}

			req := httptest.NewRequest(
				http.MethodPost,
				"/api/sendCoins",
				strings.NewReader(tc.body),
			)

			if tc.contextFromUserID {
				ctx := middleware.ContextFromUserID(
					req.Context(),
					senderID,
				)

				req = req.WithContext(ctx)
			}

			rec := httptest.NewRecorder()

			handler.Send(rec, req)

			require.Equal(
				t,
				tc.expectedStatus,
				rec.Code,
			)
		})
	}
}
