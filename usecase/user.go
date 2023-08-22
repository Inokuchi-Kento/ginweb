package usecase

import (
	"context"
	"fmt"
	"ginweb/adapter/repository"
	"ginweb/domain"
	"log"
)

type UserUseCase interface {
	CreateUser(context.Context, *domain.User) (domain.User, error)
	GetUsers(context.Context) ([]domain.ResultUser, error)
	GetUserByID(context.Context, int) (domain.ResultUser, error)
	UpdateUser(ctx context.Context, u *domain.User, id int) (domain.User, error)
	DeleteUserByID(ctx context.Context, id int) error
}

type userUseCase struct {
	UserRepository repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{UserRepository: r}
}

func (usecase *userUseCase) CreateUser(ctx context.Context, u *domain.User) (domain.User, error) {
	log.Println("usecase_u: ", u)

	user, err := usecase.UserRepository.CreateUser(ctx, u)
	if err != nil {
		err = fmt.Errorf("[usecase.CreateUser] failed: %w ", err)
	}
	return user, err
}

func (usecase *userUseCase) GetUsers(ctx context.Context) ([]domain.ResultUser, error) {
	users, err := usecase.UserRepository.GetUsers(ctx)
	if err != nil {
		err = fmt.Errorf("[usecase.GetUsers] failed: %w ", err)
	}
	return users, err
}

func (usecase *userUseCase) GetUserByID(ctx context.Context, id int) (domain.ResultUser, error) {
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
