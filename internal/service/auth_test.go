package service

import (
	"avito/internal/model"
	mock_service "avito/internal/service/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

var (
	errWithUpdateEmail = errors.New("error with update email")
	errWithCreateUser  = errors.New("error with create user")
	errWithGetUser     = errors.New("error with get user")
)

func TestService_Auth(t *testing.T) {
	password := "password"

	jwtSecret := "my-secret-key"

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	type mockBehavior func(s *mock_service.MockUserRepository)

	testUser := model.User{
		ID:           1,
		Username:     "muhammad",
		Email:        "muahammad@mail.ru",
		HashPassword: string(hashPassword),
		Coins:        1000,
	}

	testCases := []struct {
		name     string
		username string
		password string
		email    string

		mockBehavior mockBehavior

		expectedErr error
	}{
		{
			name:     "user is existing OK",
			username: testUser.Username,
			password: password,
			email:    testUser.Email,

			mockBehavior: func(s *mock_service.MockUserRepository) {
				s.EXPECT().GetByUsername(gomock.Any(), testUser.Username).Return(testUser, nil)
			},
			expectedErr: nil,
		},

		{
			name:     "user create OK",
			username: testUser.Username,
			password: password,
			email:    testUser.Email,

			mockBehavior: func(s *mock_service.MockUserRepository) {
				s.EXPECT().GetByUsername(gomock.Any(), testUser.Username).Return(model.User{}, pgx.ErrNoRows)

				s.EXPECT().Create(gomock.Any(), testUser.Username, testUser.Email, gomock.Any()).Return(testUser, nil)
			},
			expectedErr: nil,
		},
		{
			name:     "update email OK",
			username: testUser.Username,
			password: password,
			email:    "new@mail.ru",

			mockBehavior: func(s *mock_service.MockUserRepository) {
				s.EXPECT().GetByUsername(gomock.Any(), testUser.Username).Return(testUser, nil)
				s.EXPECT().UpdateEmail(gomock.Any(), testUser.ID, "new@mail.ru").Return(nil)
			},
			expectedErr: nil,
		},

		{
			name:     "wrong password",
			username: testUser.Username,
			password: "12345",
			email:    testUser.Email,

			mockBehavior: func(s *mock_service.MockUserRepository) {
				s.EXPECT().GetByUsername(gomock.Any(), testUser.Username).Return(testUser, nil)
			},
			expectedErr: ErrInvalidPassword,
		},
		{
			name:     "empty password",
			username: testUser.Username,
			password: "",
			email:    testUser.Email,

			mockBehavior: func(s *mock_service.MockUserRepository) {
			},
			expectedErr: ErrInvalidPassword,
		},

		{
			name:     "update email ERR",
			username: testUser.Username,
			password: password,
			email:    "new@mail.ru",

			mockBehavior: func(s *mock_service.MockUserRepository) {
				s.EXPECT().GetByUsername(gomock.Any(), testUser.Username).Return(testUser, nil)
				s.EXPECT().UpdateEmail(gomock.Any(), testUser.ID, "new@mail.ru").Return(errWithUpdateEmail)
			},
			expectedErr: errWithUpdateEmail,
		},
		{
			name:     "create user error",
			username: testUser.Username,
			password: password,
			email:    testUser.Email,

			mockBehavior: func(s *mock_service.MockUserRepository) {
				s.EXPECT().GetByUsername(gomock.Any(), testUser.Username).Return(model.User{}, pgx.ErrNoRows)

				s.EXPECT().Create(gomock.Any(), testUser.Username, testUser.Email, gomock.Any()).Return(model.User{}, errWithCreateUser)
			},
			expectedErr: errWithCreateUser,
		},
		{
			name:     "get user error",
			username: testUser.Username,
			password: password,
			email:    testUser.Email,

			mockBehavior: func(s *mock_service.MockUserRepository) {
				s.EXPECT().GetByUsername(gomock.Any(), testUser.Username).Return(model.User{}, errWithGetUser)
			},
			expectedErr: errWithGetUser,
		},
		{
			name:     "empty username",
			username: "",
			password: password,
			email:    testUser.Email,

			mockBehavior: func(s *mock_service.MockUserRepository) {

			},
			expectedErr: ErrEmptyUsername,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := mock_service.NewMockUserRepository(ctrl)

			tc.mockBehavior(userRepo)

			service := &Service{
				UserRepository: userRepo,
				jwtSecret:      jwtSecret,
			}

			token, err := service.Auth(context.Background(), tc.username, tc.password, tc.email)

			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
				require.Empty(t, token)
				return
			}

			require.NoError(t, err)
			require.NotEmpty(t, token)
		})
	}
}
