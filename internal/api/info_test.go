package api

import (
	"avito/internal/api/middleware"
	mock_api "avito/internal/api/mocks"
	"avito/internal/model"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

var errWithInfo = errors.New("error with send")

func TestApi_Info(t *testing.T) {
	logger, _ := test.NewNullLogger()
	type mockBehaviour func(serviceMock *mock_api.MockService)

	userID := int64(1)

	testCases := []struct {
		name              string
		userID            int64
		contextFromUserID bool

		mockBehaviour mockBehaviour

		expectedStatus int
	}{
		{
			name:              "info api OK",
			userID:            userID,
			contextFromUserID: true,

			mockBehaviour: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().Info(
					gomock.Any(),
					userID,
				)
			},

			expectedStatus: http.StatusOK,
		},

		{
			name:              "info context ERROR",
			userID:            userID,
			contextFromUserID: false,

			mockBehaviour: func(serviceMock *mock_api.MockService) {
			},

			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:              "info api ERR",
			userID:            userID,
			contextFromUserID: true,

			mockBehaviour: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().Info(
					gomock.Any(),
					userID,
				).Return(model.Info{}, errWithInfo)
			},

			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			serviceMock := mock_api.NewMockService(
				ctrl,
			)

			tc.mockBehaviour(
				serviceMock,
			)

			handler := New(
				serviceMock,
				logger,
			)

			req := httptest.NewRequestWithContext(
				context.Background(),
				http.MethodGet,
				"/api/info",
				nil,
			)

			if tc.contextFromUserID {
				ctx := middleware.ContextFromUserID(
					req.Context(),
					userID,
				)

				req = req.WithContext(
					ctx,
				)
			}

			rec := httptest.NewRecorder()

			handler.Info(
				rec,
				req,
			)

			require.Equal(
				t,
				tc.expectedStatus,
				rec.Code,
			)
		})
	}
}
