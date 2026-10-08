package api

import (
	"avito/internal/api/middleware"
	mock_api "avito/internal/api/mocks"
	"avito/internal/service"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

var errService = errors.New("error with service")

func TestApi_buy(t *testing.T) {
	type mockBehavior func(serviceMock *mock_api.MockService)
	logger, _ := test.NewNullLogger()

	itemName := "book"
	userID := int64(1)

	testCases := []struct {
		name     string
		itemName string
		userID   int64

		contextFromUserID bool

		mockBehavior mockBehavior

		expectedStatus int
	}{
		{
			name:     "buy api OK",
			userID:   userID,
			itemName: itemName,

			contextFromUserID: true,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().Buy(
					gomock.Any(),
					userID,
					itemName,
				).Return(nil)
			},

			expectedStatus: http.StatusOK,
		},

		{
			name:     "empty item",
			userID:   userID,
			itemName: itemName,

			contextFromUserID: true,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().Buy(
					gomock.Any(),
					userID,
					itemName,
				).Return(service.ErrEmptyItem)
			},

			expectedStatus: http.StatusBadRequest,
		},

		{
			name:              "not enough coins",
			contextFromUserID: true,
			itemName:          itemName,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().
					Buy(
						gomock.Any(),
						userID,
						itemName,
					).
					Return(service.ErrNotEnoughCoins)
			},

			expectedStatus: http.StatusBadRequest,
		},
		{
			name:              "service buy error",
			contextFromUserID: true,
			itemName:          itemName,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().
					Buy(
						gomock.Any(),
						userID,
						itemName,
					).
					Return(errService)
			},

			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:              "user ID not in context",
			contextFromUserID: false,
			itemName:          itemName,

			mockBehavior: func(serviceMock *mock_api.MockService) {
			},

			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			serviceMock := mock_api.NewMockService(ctrl)

			tc.mockBehavior(
				serviceMock,
			)

			handler := New(
				serviceMock,
				logger,
			)

			req := httptest.NewRequest(
				http.MethodGet,
				"/api/buy/"+tc.itemName,
				nil,
			)

			req.SetPathValue(
				"item",
				tc.itemName,
			)

			if tc.contextFromUserID {
				ctx := middleware.ContextFromUserID(
					req.Context(),
					userID,
				)

				req = req.WithContext(ctx)
			}

			rec := httptest.NewRecorder()

			handler.Buy(rec, req)

			require.Equal(
				t,
				tc.expectedStatus,
				rec.Code,
			)
		})
	}
}
