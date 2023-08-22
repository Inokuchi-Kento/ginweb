package usecase

import (
	"context"
	"fmt"
	"ginweb/adapter/repository"
	"ginweb/domain"
)

type GroupUseCase interface {
	CreateGroup(context.Context, *domain.Group) (domain.Group, error)
	UpdateGroupName(context.Context, *domain.Group, int) (domain.Group, error)
}

type groupUseCase struct {
	GroupRepository repository.GroupRepository
}

func NewGroupUseCase(gr repository.GroupRepository) GroupUseCase {
	return &groupUseCase{GroupRepository: gr}
}

func (usecase *groupUseCase) CreateGroup(ctx context.Context, g *domain.Group) (domain.Group, error) {
	group, err := usecase.GroupRepository.CreateGroup(ctx, g)
	if err != nil {
		err = fmt.Errorf("[usecase.CreateGroup] failed: %w ", err)
	}
	return group, err
}

func (usecase *groupUseCase) UpdateGroupName(ctx context.Context, g *domain.Group, id int) (domain.Group, error) {
	group, err := usecase.GroupRepository.UpdateGroupName(ctx, g, id)
	if err != nil {
		err = fmt.Errorf("[usecase.UpdateGroupName] failed: %w ", err)
	}
	return group, err
}
