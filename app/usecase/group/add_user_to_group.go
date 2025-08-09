package group

import (
	"context"

	groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
)

type AddUserToGroupUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewAddUserToGroupUseCase(
	groupRepo groupDomain.GroupRepository,
) *AddUserToGroupUseCase {
	return &AddUserToGroupUseCase{
		groupRepo: groupRepo,
	}
}

type AddUserToGroupUseCaseDto struct {
	UserID  string
	GroupID string
}

func (uc *AddUserToGroupUseCase) Run(ctx context.Context, dto AddUserToGroupUseCaseDto) (*groupDomain.Group, error) {
	err := uc.groupRepo.AddUserToGroup(ctx, dto.GroupID, dto.UserID)
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
