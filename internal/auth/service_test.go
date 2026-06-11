package auth

import (
	"errors"
	"os"
	"testing"

	"akubisa/pkg/logger"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// MockRepository is a mock implementation of the Repository interface
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateUser(user User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockRepository) FindByEmail(email string) (*User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func TestMain(m *testing.M) {
	logger.Init("test")
	os.Exit(m.Run())
}

func TestService_Register(t *testing.T) {
	tests := []struct {
		name          string
		req           RegisterRequest
		mockBehavior  func(m *MockRepository)
		expectedError bool
	}{
		{
			name: "Success",
			req: RegisterRequest{
				FullName: "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			mockBehavior: func(m *MockRepository) {
				m.On("CreateUser", mock.MatchedBy(func(user User) bool {
					return user.FullName == "John Doe" && user.Email == "john@example.com"
				})).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Repository Error",
			req: RegisterRequest{
				FullName: "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			mockBehavior: func(m *MockRepository) {
				m.On("CreateUser", mock.Anything).Return(errors.New("database error"))
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockBehavior(mockRepo)

			service := NewService(mockRepo, "secret")
			err := service.Register(tt.req)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_Login(t *testing.T) {
	secret := "test-secret"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), 12)
	userID := uuid.New()

	tests := []struct {
		name          string
		req           LoginRequest
		mockBehavior  func(m *MockRepository)
		expectedError string
	}{
		{
			name: "Success",
			req: LoginRequest{
				Email:    "john@example.com",
				Password: "password123",
			},
			mockBehavior: func(m *MockRepository) {
				m.On("FindByEmail", "john@example.com").Return(&User{
					ID:           userID,
					Email:        "john@example.com",
					PasswordHash: string(hashedPassword),
				}, nil)
			},
			expectedError: "",
		},
		{
			name: "User Not Found",
			req: LoginRequest{
				Email:    "notfound@example.com",
				Password: "password123",
			},
			mockBehavior: func(m *MockRepository) {
				m.On("FindByEmail", "notfound@example.com").Return(nil, errors.New("not found"))
			},
			expectedError: "invalid credentials",
		},
		{
			name: "Invalid Password",
			req: LoginRequest{
				Email:    "john@example.com",
				Password: "wrongpassword",
			},
			mockBehavior: func(m *MockRepository) {
				m.On("FindByEmail", "john@example.com").Return(&User{
					ID:           userID,
					Email:        "john@example.com",
					PasswordHash: string(hashedPassword),
				}, nil)
			},
			expectedError: "invalid credentials",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockBehavior(mockRepo)

			service := NewService(mockRepo, secret)
			token, err := service.Login(tt.req)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
