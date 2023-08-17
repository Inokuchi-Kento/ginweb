package usecase

import (
	"context"
	"ginweb/domain"
	"log"
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

func (m *MockUserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, u *domain.User, id int) (domain.User, error) {
	args := m.Called(ctx, u, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUserByID(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestUserUseCase_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsesase := NewUserUseCase(mockRepo)

	ctx := context.Background()
	inputUser := domain.User{
		Name: "test",
		Age:  20,
	}

	createdUser := domain.User{
		Name: "test",
		Age:  20,
	}

	mockRepo.On("CreateUser", ctx, &inputUser).Return(createdUser, nil)

	resultUser, err := userUsesase.CreateUser(ctx, &inputUser)
	log.Println(resultUser)

	mockRepo.AssertCalled(t, "CreateUser", ctx, &inputUser)

	assert.NoError(t, err)
	assert.Equal(t, createdUser, resultUser)
}
