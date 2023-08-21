package usecase

import (
	"context"
	"ginweb/adapter/repository/mock"
	"ginweb/domain"
	"ginweb/domain/testdata"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUseCase_CreateUser(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)
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

func TestUserUseCase_GetUsers(t *testing.T) {
	mockRepo := new(mock.MockUserRepository)
	userUsesase := NewUserUseCase(mockRepo)

	ctx := context.Background()

	getUsers := testdata.Users

	mockRepo.On("GetUsers", ctx).Return(getUsers, nil)

	resultUsers, err := userUsesase.GetUsers(ctx)
	log.Println(resultUsers)

	mockRepo.AssertCalled(t, "GetUsers", ctx)

	assert.NoError(t, err)
	assert.Equal(t, getUsers, resultUsers)
}
