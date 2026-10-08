package api

import (
	mock_api "avito/internal/api/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"avito/internal/service"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func TestHandler_Auth(t *testing.T) {
	logger, _ := test.NewNullLogger()
	token := "test-token"

	errService := errors.New("service error")

	type mockBehavior func(
		serviceMock *mock_api.MockService,
	)

	testCases := []struct {
		name string
		body string

		mockBehavior mockBehavior

		expectedStatus int
	}{
		{
			name: "auth OK",
			body: `{
				"username":"muhammad",
				"password":"wrong-password",
				"email":"muhammad@mail.ru"
			}`,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().
					Auth(
						gomock.Any(),
						"muhammad",
						"wrong-password",
						"muhammad@mail.ru",
					).
					Return(token, nil)
			},

			expectedStatus: http.StatusOK,
		},

		{
			name: "invalid json",
			body: `{`,

			mockBehavior: func(serviceMock *mock_api.MockService) {
			},

			expectedStatus: http.StatusBadRequest,
		},

		{
			name: "invalid email",
			body: `{
				"username":"muhammad",
				"password":"password",
				"email":"muh"
			}`,

			mockBehavior: func(serviceMock *mock_api.MockService) {
			},

			expectedStatus: http.StatusBadRequest,
		},

		{
			name: "invalid password",
			body: `{
				"username":"muhammad",
				"password":"password",
				"email":"muhammad@mail.ru"
			}`,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().
					Auth(
						gomock.Any(),
						"muhammad",
						"password",
						"muhammad@mail.ru",
					).
					Return("", service.ErrInvalidPassword)
			},

			expectedStatus: http.StatusBadRequest,
		},

		{
			name: "service error",
			body: `{
				"username":"muhammad",
				"password":"password",
				"email":"muhammad@mail.ru"
			}`,

			mockBehavior: func(serviceMock *mock_api.MockService) {
				serviceMock.EXPECT().
					Auth(
						gomock.Any(),
						"muhammad",
						"password",
						"muhammad@mail.ru",
					).
					Return("", errService)
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
				"/api/auth",
				strings.NewReader(tc.body),
			)

			rec := httptest.NewRecorder()

			handler.Auth(rec, req)

			require.Equal(
				t,
				tc.expectedStatus,
				rec.Code,
			)
		})
	}
}
