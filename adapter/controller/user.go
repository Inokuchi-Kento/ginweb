package controller

import (
	"ginweb/adapter/repository"
	"ginweb/domain"
	"ginweb/ent"
	"ginweb/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UseCase usecase.UserUseCase
}

func NewUserController(db *ent.Client) *userController {
	return &userController{
		UseCase: usecase.NewUserUseCase(repository.NewUserRepository(db)),
	}
}

func (controller *userController) GetUsers(ctx *gin.Context) {
	users, err := controller.UseCase.GetUsers(ctx)
	if checkError(ctx, err) {
		return
	}
	responseOK(ctx, users)
}

func (controller *userController) GetUserByID(ctx *gin.Context) {
	id, ok := parsePathInt(ctx, "id")
	if !ok {
		return
	}

	user, err := controller.UseCase.GetUserByID(ctx, id)
	if err != nil {
		return
	}
	responseOK(ctx, user)
}

func (controller *userController) CreateUser(ctx *gin.Context) {
	u := &domain.User{}
	ok := bindJSON(ctx, u)
	if !ok {
		return
	}

	log.Println(u)

	user, err := controller.UseCase.CreateUser(ctx, u)
	if checkError(ctx, err) {
		return
	}
	responseCreated(ctx, user)
}

func (controller *userController) UpdateUser(ctx *gin.Context) {
	id, ok := parsePathInt(ctx, "id")
	if !ok {
		return
	}
	u := &domain.User{}
	ok = bindJSON(ctx, u)
	if !ok {
		return
	}
	user, err := controller.UseCase.UpdateUser(ctx, u, id)
	if checkError(ctx, err) {
		return
	}
	responseOK(ctx, user)
}

func (controller *userController) DeleteUser(ctx *gin.Context) {
	id, ok := parsePathInt(ctx, "id")
	if !ok {
		return
	}

	err := controller.UseCase.DeleteUserByID(ctx, id)
	if checkError(ctx, err) {
		return
	}
	responseNoContent(ctx)
}

// upsert
func (controller *userController) UpsertUser(ctx *gin.Context) {
	u := &domain.User{}
	ok := bindJSON(ctx, u)
	if !ok {
		return
	}

	id, err := controller.UseCase.UpsertUser(ctx, u)
	if checkError(ctx, err) {
		return
	}
	responseOK(ctx, id)
}
