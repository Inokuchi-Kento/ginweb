package usecase

import (
	"context"
	"ginweb/domain"

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
