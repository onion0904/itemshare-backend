package group

import (
	"context"

	groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
)

type RemoveUserFromGroupUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewRemoveUserToGroupUseCase(
	groupRepo groupDomain.GroupRepository,
) *RemoveUserFromGroupUseCase {
	return &RemoveUserFromGroupUseCase{
		groupRepo: groupRepo,
	}
}

type RemoveUserFromGroupUseCaseDto struct {
	UserID  string
	GroupID string
}

func (uc *RemoveUserFromGroupUseCase) Run(ctx context.Context, dto RemoveUserFromGroupUseCaseDto) (*groupDomain.Group, error) {
	err := uc.groupRepo.RemoveUserFromGroup(ctx, dto.GroupID, dto.UserID)
	if err != nil {
		return nil, err
	}
	group, err := uc.groupRepo.FindGroupByID(ctx, dto.GroupID)
	if err != nil {
		return nil, err
	}
	group, err = groupDomain.Reconstruct(group.ID(), group.Name(), group.UserIDs(), group.EventIDs())
	if err != nil {
		return nil, err
	}
	return group, nil
}
