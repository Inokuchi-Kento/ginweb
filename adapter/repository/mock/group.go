package mock

import (
	"context"
	"ginweb/domain"

	"github.com/stretchr/testify/mock"
)

type MockGroupRepository struct {
	mock.Mock
}

func (m *MockGroupRepository) CreateGroup(ctx context.Context, g *domain.Group) (domain.Group, error) {
	args := m.Called(ctx, g)
	return args.Get(0).(domain.Group), args.Error(1)
}

func (m *MockGroupRepository) UpdateGroupName(ctx context.Context, g *domain.Group, id int) (domain.Group, error) {
	args := m.Called(ctx, g, id)
	return args.Get(0).(domain.Group), args.Error(1)
}

func (m *MockGroupRepository) GetGroups(ctx context.Context) ([]domain.Group, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Group), args.Error(1)
}
