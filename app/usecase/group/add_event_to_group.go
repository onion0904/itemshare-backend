package group

import (
	"context"
	groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
)

type AddEventToGroupUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewAddEventToGroupUseCase(
	groupRepo groupDomain.GroupRepository,
) *AddEventToGroupUseCase {
	return &AddEventToGroupUseCase{
		groupRepo: groupRepo,
	}
}

type AddEventToGroupUseCaseDto struct {
	EventID string
	GroupID string
}

//
func (uc *AddEventToGroupUseCase) Run(ctx context.Context, dto AddEventToGroupUseCaseDto) (*groupDomain.Group,error) {
	err := uc.groupRepo.AddEventToGroup(ctx, dto.GroupID, dto.EventID)
	if err != nil {
		return nil,err
	}
	group, err := uc.groupRepo.FindGroup(ctx, dto.GroupID)
	if err != nil {
        return nil,err
    }
	group, err = groupDomain.Reconstruct(group.ID(), group.Name(), group.UserIDs(),group.EventIDs() ,group.Icon())
	if err != nil {
		return nil,err
	}
	return group,nil
}