package group

import (
	"context"
	groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
)

type DeleteUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewDeleteUseCase(
	groupRepo groupDomain.GroupRepository,
) *DeleteUseCase {
	return &DeleteUseCase{
		groupRepo: groupRepo,
	}
}

func (uc *DeleteUseCase) Run(ctx context.Context, groupID string) error {
	err := uc.groupRepo.Delete(ctx, groupID)
	if err != nil {
		return err
	}
	return nil
}
