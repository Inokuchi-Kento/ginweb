package usecase

import (
	"context"
	"fmt"
	"ginweb/adapter/repository/mock"
	"ginweb/domain/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupUseCase_CreateGroup(t *testing.T) {
	mockRepo := new(mock.MockGroupRepository)
	groupUseCase := NewGroupUseCase(mockRepo)

	ctx := context.Background()
	for _, inputGroup := range testdata.Groups {
		// Capture the inputGroup in the loop scope
		inputGroup := inputGroup

		t.Run(fmt.Sprintf("Creating group with ID %d", inputGroup.ID), func(t *testing.T) {
			createdData := inputGroup

			mockRepo.On("CreateGroup", ctx, &inputGroup).Return(createdData, nil)
			resultGroup, err := groupUseCase.CreateGroup(ctx, &inputGroup)

			mockRepo.AssertCalled(t, "CreateGroup", ctx, &inputGroup)
			assert.NoError(t, err)
			assert.Equal(t, createdData, resultGroup)
		})
	}
}

func TestGroupUseCase_UpdateGroup(t *testing.T) {
	mockRepo := new(mock.MockGroupRepository)
	groupUseCase := NewGroupUseCase(mockRepo)

	ctx := context.Background()
	for _, inputGroup := range testdata.Groups {
		// Capture the inputGroup in the loop scope
		inputGroup := inputGroup

		t.Run(fmt.Sprintf("Updating group with ID %d", inputGroup.ID), func(t *testing.T) {
			updatedData := inputGroup

			mockRepo.On("UpdateGroupName", ctx, &inputGroup, inputGroup.ID).Return(updatedData, nil)
			resultGroup, err := groupUseCase.UpdateGroupName(ctx, &inputGroup, inputGroup.ID)

			mockRepo.AssertCalled(t, "UpdateGroupName", ctx, &inputGroup, inputGroup.ID)
			assert.NoError(t, err)
			assert.Equal(t, updatedData, resultGroup)
		})
	}
}

// GetGroupsをテストする
// テストケースはtestdata/groups.goに定義
// テストケースの数だけループしてテストを実行する
func TestGroupUseCase_GetGroups(t *testing.T) {
	mockRepo := new(mock.MockGroupRepository)
	groupUseCase := NewGroupUseCase(mockRepo)

	ctx := context.Background()
	for _, inputGroup := range testdata.Groups {
		// Capture the inputGroup in the loop scope
		inputGroup := inputGroup

		t.Run(fmt.Sprintf("Getting group with ID %d", inputGroup.ID), func(t *testing.T) {
			// テストデータを設定
			testGroups := testdata.Groups

			mockRepo.On("GetGroups", ctx).Return(testGroups, nil)
			resultGroups, err := groupUseCase.GetGroups(ctx)

			mockRepo.AssertCalled(t, "GetGroups", ctx)
			assert.NoError(t, err)
			assert.Equal(t, testGroups, resultGroups)
		})
	}
}
