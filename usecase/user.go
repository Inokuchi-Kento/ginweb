package usecase

import (
	"context"
	"fmt"
	"ginweb/adapter/repository"
	"ginweb/domain"
)

type UserUseCase interface {
	CreateUser(context.Context, *domain.User) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id int) (domain.User, error)
	UpdateUser(ctx context.Context, u *domain.User, id int) (domain.User, error)
	DeleteUserByID(ctx context.Context, id int) error
	UpsertUser(ctx context.Context, u *domain.User) (int, error)
}

type userUseCase struct {
	UserRepository repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{UserRepository: r}
}

func (usecase *userUseCase) CreateUser(ctx context.Context, u *domain.User) (domain.User, error) {
	user, err := usecase.UserRepository.CreateUser(ctx, u)
	if err != nil {
		err = fmt.Errorf("[usecase.CreateUser] failed: %w ", err)
	}
	return user, err
}

func (usecase *userUseCase) GetUsers(ctx context.Context) ([]domain.User, error) {
	users, err := usecase.UserRepository.GetUsers(ctx)
	if err != nil {
		err = fmt.Errorf("[usecase.GetUsers] failed: %w ", err)
	}
	return users, err
}

func (usecase *userUseCase) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	user, err := usecase.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("[usecase.GetUserByID] failed: %w ", err)
	}
	return user, err
}

func (usecase *userUseCase) UpdateUser(ctx context.Context, u *domain.User, id int) (domain.User, error) {
	user, err := usecase.UserRepository.UpdateUser(ctx, u, id)
	if err != nil {
		err = fmt.Errorf("[usecase.UpdateUser] failed: %w ", err)
	}
	return user, err
}

func (usecase *userUseCase) DeleteUserByID(ctx context.Context, id int) error {
	err := usecase.UserRepository.DeleteUserByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("[usecase.DeleteUserByID] failed: %w ", err)
	}
	return err
}

// uppsert
func (usecase *userUseCase) UpsertUser(ctx context.Context, u *domain.User) (int, error) {
	affected, err := usecase.UserRepository.UpsertUser(ctx, u)
	if err != nil {
		err = fmt.Errorf("[usecase.UpsertUser] failed: %w ", err)
	}
	return affected, err
}
