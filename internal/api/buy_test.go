package api

// import (
// 	"avito/internal/api/middleware"
// 	mock_api "avito/internal/api/mocks"
// 	"context"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"path"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/sirupsen/logrus/hooks/test"
// 	"github.com/stretchr/testify/require"
// )

// func TestHandler_Buy(t *testing.T) {
// 	var (
// 		errWithBuy = errors.New("error with buy")
// 	)

// 	type mockBehavior func(buyService *mock_api.MockBuyService)
// 	logger, _ := test.NewNullLogger()

// 	userID := int64(1)

// 	testCases := []struct {
// 		name   string
// 		path   string
// 		userID int64

// 		mockBehavior mockBehavior

// 		expectedCode int
// 	}{
// 		{
// 			name:   "buy OK",
// 			path:   "/api/buy/book",
// 			userID: userID,

// 			mockBehavior: func(buyService *mock_buy.MockBuyService) {
// 				buyService.EXPECT().Buy(
// 					gomock.Any(),
// 					userID,
// 					"book",
// 				)
// 			},
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:   "buy ERR",
// 			path:   "/api/buy/book",
// 			userID: userID,

// 			mockBehavior: func(buyService *mock_buy.MockBuyService) {
// 				buyService.EXPECT().Buy(
// 					gomock.Any(),
// 					userID,
// 					"book",
// 				).Return(errWithBuy)
// 			},
// 			expectedCode: http.StatusInternalServerError,
// 		},
// 		{
// 			name:   "invalid item name",
// 			path:   "/api/buy/book",
// 			userID: userID,

// 			mockBehavior: func(buyService *mock_buy.MockBuyService) {
// 				buyService.EXPECT().Buy(
// 					gomock.Any(),
// 					userID,
// 					"book",
// 				).Return(buyservice.ErrWithItemName)
// 			},
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:   "not enough coins",
// 			path:   "/api/buy/book",
// 			userID: userID,

// 			mockBehavior: func(buyService *mock_buy.MockBuyService) {
// 				buyService.EXPECT().Buy(
// 					gomock.Any(),
// 					userID,
// 					"book",
// 				).Return(buyservice.ErrWithEnoughCoins)
// 			},

// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name: "error with context",
// 			path: "/api/buy/book",

// 			mockBehavior: func(buyService *mock_buy.MockBuyService) {

// 			},

// 			expectedCode: http.StatusInternalServerError,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)

// 			service := mock_buy.NewMockBuyService(ctrl)
// 			tc.mockBehavior(service)

// 			handler := New(
// 				service,
// 				logger,
// 			)

// 			req := httptest.NewRequest(
// 				http.MethodPost,
// 				tc.path,
// 				nil,
// 			)

// 			itemName := path.Base(tc.path)
// 			req.SetPathValue("item", itemName)

// 			var ctx context.Context

// 			if tc.name == "error with context" {
// 				ctx = req.Context()
// 			} else {
// 				ctx = middleware.ContextFromUserID(
// 					req.Context(),
// 					tc.userID,
// 				)
// 			}

// 			req = req.WithContext(ctx)

// 			rec := httptest.NewRecorder()

// 			handler.Buy(rec, req)

// 			require.Equal(
// 				t,
// 				tc.expectedCode,
// 				rec.Code,
// 			)
// 		})
// 	}
// }
