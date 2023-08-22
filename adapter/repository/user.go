package repository

import (
	"context"
	"database/sql"
	"fmt"
	"ginweb/domain"
	"ginweb/ent"
	"ginweb/ent/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u *domain.User) (user domain.User, err error)
	GetUsers(ctx context.Context) (users []domain.ResultUser, err error)
	GetUserByID(ctx context.Context, id int) (user domain.ResultUser, err error)
	UpdateUser(ctx context.Context, u *domain.User, id int) (user domain.User, err error)
	DeleteUserByID(ctx context.Context, id int) error
}

type userRepository struct {
	DB *ent.Client
}

func NewUserRepository(db *ent.Client) (repo UserRepository) {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) GetUsers(ctx context.Context) (users []domain.ResultUser, err error) {
	getUser, err := ur.DB.User.Query().
		Order(ent.Asc(user.FieldID)).
		WithGroup().
		All(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.GetUsers] failed to find all users: %w", err)
		return
	}
	if len(getUser) == 0 {
		err = fmt.Errorf("[repository.UserRepository.GetUsers] record not found: %w", sql.ErrNoRows)
		return
	}
	for _, item := range getUser {
		user := domain.ResultUser{
			ID:   item.ID,
			Name: item.Name,
			Age:  item.Age,
			Group: domain.Group{
				ID:   item.Edges.Group.ID,
				Name: item.Edges.Group.Name,
			},
		}
		users = append(users, user)
	}
	return
}

func (ur *userRepository) GetUserByID(ctx context.Context, id int) (user domain.ResultUser, err error) {
	getUser, err := ur.DB.User.Get(ctx, id)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.GetUserByID] failed to find user: %w", err)
		return user, err
	}
	user = domain.ResultUser{
		ID:   getUser.ID,
		Name: getUser.Name,
		Age:  getUser.Age,
		Group: domain.Group{
			ID:   getUser.Edges.Group.ID,
			Name: getUser.Edges.Group.Name,
		},
	}
	return user, nil
}

func (ur *userRepository) CreateUser(ctx context.Context, u *domain.User) (user domain.User, err error) {
	createUserRow, err := ur.DB.User.Create().
		SetName(u.Name).
		SetAge(u.Age).
		SetGroupID(u.Group.ID).
		Save(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.CreateUser] failed: user = %+v, error = %w ", user, err)
		return user, err
	}
	user = domain.User{
		ID:   createUserRow.ID,
		Name: createUserRow.Name,
		Age:  createUserRow.Age,
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, u *domain.User, id int) (user domain.User, err error) {
	updateUser, err := ur.DB.User.UpdateOneID(id).
		SetName(u.Name).
		SetAge(u.Age).
		Save(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.UpdateUser] failed: user = %+v, error = %w ", user, err)
		return user, err
	}
	user = domain.User{
		ID:   updateUser.ID,
		Name: updateUser.Name,
		Age:  updateUser.Age,
	}
	return user, nil
}

func (ur *userRepository) DeleteUserByID(ctx context.Context, id int) error {
	err := ur.DB.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.UserRepository.DeleteUserByID] failed: id = %d, error = %w ", id, err)
	}
	return err
}
