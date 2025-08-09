package group

import (
	"context"

	groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
)

type UpdateUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewUpdateUseCase(
	groupRepo groupDomain.GroupRepository,
) *UpdateUseCase {
	return &UpdateUseCase{
		groupRepo: groupRepo,
	}
}

type UpdateUseCaseDto struct {
	Name    string
	UsersID []string
}

func (uc *UpdateUseCase) Run(ctx context.Context, groupID string, dto UpdateUseCaseDto) (*groupDomain.Group, error) {
	// dtoからuserへ変換
	group, err := uc.groupRepo.FindGroupByID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	ngroup, err := groupDomain.Reconstruct(groupID, dto.Name, group.UserIDs(), group.EventIDs())
	if err != nil {
		return nil, err
	}
	err = uc.groupRepo.Update(ctx, ngroup)
	if err != nil {
		return nil, err
	}
	return uc.groupRepo.FindGroupByID(ctx, ngroup.ID())
}
