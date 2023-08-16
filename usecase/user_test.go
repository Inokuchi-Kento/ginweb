package usecase

import (
	"context"
	"ginweb/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, u *domain.User) (domain.User, error) {
	args := m.Called(ctx, u)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUserByID(ctx context.Context, userID int) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	u := domain.User{
		Name: "test",
		Age:  20,
	}
	mockRepo := new(MockUserRepository)
	mockRepo.On("CreateUser", ctx, &u).Return(u, nil)

	uc := NewUserUseCase(mockRepo)

	user, err := uc.CreateUser(ctx, &u)
	assert.NoError(t, err)
	assert.Equal(t, u, user)
}
