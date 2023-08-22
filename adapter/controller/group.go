package controller

import (
	"ginweb/adapter/repository"
	"ginweb/domain"
	"ginweb/ent"
	"ginweb/usecase"

	"github.com/gin-gonic/gin"
)

type groupController struct {
	UseCase usecase.GroupUseCase
}

func NewGroupController(db *ent.Client) *groupController {
	return &groupController{UseCase: usecase.NewGroupUseCase(repository.NewGroupRepository(db))}
}

func (controller *groupController) CreateGroup(ctx *gin.Context) {
	g := &domain.Group{}
	ok := bindJSON(ctx, g)
	if !ok {
		return
	}

	group, err := controller.UseCase.CreateGroup(ctx, g)
	if checkError(ctx, err) {
		return
	}
	responseCreated(ctx, group)
}

func (controller *groupController) UpdateGroupName(ctx *gin.Context) {
	g := &domain.Group{}
	ok := bindJSON(ctx, g)
	if !ok {
		return
	}

	id, ok := parsePathInt(ctx, "id")
	if !ok {
		return
	}

	group, err := controller.UseCase.UpdateGroupName(ctx, g, id)
	if checkError(ctx, err) {
		return
	}
	responseOK(ctx, group)
}
