package usecase

import (
	"context"
	"fmt"
	"ginweb/adapter/repository"
	"ginweb/domain"
)

type GroupUseCase interface {
	CreateGroup(context.Context, *domain.Group) (domain.Group, error)
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
