package repository

import (
	"context"
	"fmt"
	"ginweb/domain"
	"ginweb/ent"
)

type GroupRepository interface {
	CreateGroup(ctx context.Context, g *domain.Group) (group domain.Group, err error)
	UpdateGroupName(ctx context.Context, g *domain.Group, id int) (group domain.Group, err error)
}

type groupRepository struct {
	DB *ent.Client
}

func NewGroupRepository(db *ent.Client) GroupRepository {
	return &groupRepository{
		DB: db,
	}
}

func (gr *groupRepository) CreateGroup(ctx context.Context, g *domain.Group) (group domain.Group, err error) {
	createGroupRow, err := gr.DB.Group.Create().
		SetName(g.Name).
		Save(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.GroupRepository.CreateGroup] failed to create group: %w", err)
		return group, err
	}
	group = domain.Group{
		ID:   createGroupRow.ID,
		Name: createGroupRow.Name,
	}
	return group, nil
}

// update
func (gr *groupRepository) UpdateGroupName(ctx context.Context, g *domain.Group, id int) (group domain.Group, err error) {
	updateGroupRow, err := gr.DB.Group.UpdateOneID(id).
		SetName(g.Name).
		Save(ctx)
	if err != nil {
		err = fmt.Errorf("[repository.GroupRepository.UpdateGroup] failed to update group: %w", err)
		return group, err
	}
	group = domain.Group{
		ID:   updateGroupRow.ID,
		Name: updateGroupRow.Name,
	}
	return group, nil
}
