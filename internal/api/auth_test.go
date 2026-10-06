package api

// import (
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/sirupsen/logrus/hooks/test"
// 	"github.com/stretchr/testify/require"
// )

// func TestHandler_Auth(t *testing.T) {
// 	var (
// 		errWithToken = errors.New("error with token")
// 	)

// 	type mockBehaviour func(authService *mock_auth.MockAuthService)
// 	logger, _ := test.NewNullLogger()

// 	testcases := []struct {
// 		name string
// 		body string

// 		mockBehaviour mockBehaviour

// 		expectedCode int
// 	}{
// 		{
// 			name: "auth OK",
// 			body: `{
// 			"username":"muhammad",
// 			"email":"muhammad@mail.ru",
// 			"password":"password"
// 			}`,

// 			mockBehaviour: func(authService *mock_auth.MockAuthService) {
// 				authService.EXPECT().Auth(
// 					gomock.Any(),
// 					"muhammad",
// 					"password",
// 					"muhammad@mail.ru",
// 				).Return("test-token", nil)
// 			},

// 			expectedCode: http.StatusOK,
// 		},

// 		{
// 			name: "auth ERR",
// 			body: `{
// 			"username":"muhammad",
// 			"email":"muhammad@mail.ru",
// 			"password":"password"
// 			}`,

// 			mockBehaviour: func(authService *mock_auth.MockAuthService) {
// 				authService.EXPECT().Auth(
// 					gomock.Any(),
// 					"muhammad",
// 					"password",
// 					"muhammad@mail.ru",
// 				).Return("test-token", errWithToken)
// 			},

// 			expectedCode: http.StatusInternalServerError,
// 		},

// 		{
// 			name: "decode ERR",
// 			body: `{
// 			"username":"muhammad",
// 			"email":"muhammad@mail.ru",
// 			"password":"password"
// 			`,

// 			mockBehaviour: func(authService *mock_auth.MockAuthService) {

// 			},

// 			expectedCode: http.StatusBadRequest,
// 		},

// 		{
// 			name: "invalid password",
// 			body: `{
// 			"username":"muhammad",
// 			"password":"",
// 			"email":"muhammad@mail.ru"
// 			}`,

// 			mockBehaviour: func(authService *mock_auth.MockAuthService) {

// 			},

// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name: "invalid email",
// 			body: `{
// 			"username":"muhammad",
// 			"password":"password",
// 			"email":"muhammad"
// 			}`,

// 			mockBehaviour: func(authService *mock_auth.MockAuthService) {

// 			},

// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name: "error with parse token",
// 			body: `{
// 			"username":"muhammad",
// 			"password":"password",
// 			"email":"muhammad"
// 			}`,

// 			mockBehaviour: func(authService *mock_auth.MockAuthService) {

// 			},

// 			expectedCode: http.StatusUnauthorized,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)

// 			authService := mock_auth.NewMockAuthService(ctrl)
// 			tc.mockBehaviour(authService)

// 			handler := New(
// 				authService,
// 				logger,
// 			)

// 			req := httptest.NewRequest(
// 				http.MethodPost,
// 				"/api/buy",
// 				strings.NewReader(tc.body),
// 			)

// 			rec := httptest.NewRecorder()

// 			handler.Auth(rec, req)

// 			require.Equal(
// 				t,
// 				tc.expectedCode,
// 				rec.Code,
// 			)
// 		})
// 	}
// }
