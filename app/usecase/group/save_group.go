package group

import (
	"context"

	groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
)

type SaveUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewSaveUseCase(
	groupRepo groupDomain.GroupRepository,
) *SaveUseCase {
	return &SaveUseCase{
		groupRepo: groupRepo,
	}
}

type SaveUseCaseDto struct {
	Name    string
	UsersID []string
}

func (uc *SaveUseCase) Run(ctx context.Context, dto SaveUseCaseDto) (*groupDomain.Group, error) {
	// dtoからuserへ変換
	group, err := groupDomain.NewGroup(dto.Name, dto.UsersID)
	if err != nil {
		return nil, err
	}
	err = uc.groupRepo.Save(ctx, group)
	if err != nil {
		return nil, err
	}
	return uc.groupRepo.FindGroupByID(ctx, group.ID())
}
