package group

import (
	"context"
	"time"

	groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
)

type FindGroupUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewFindGroupUseCase(
	groupRepo groupDomain.GroupRepository,
) *FindGroupUseCase {
	return &FindGroupUseCase{
		groupRepo: groupRepo,
	}
}

type FindGroupUseCaseDto struct {
	ID        string
	Name      string
	UserIDs   []string
	EventIDs  []string
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (uc *FindGroupUseCase) Run(ctx context.Context, groupID string) (*FindGroupUseCaseDto, error) {
	group, err := uc.groupRepo.FindGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return &FindGroupUseCaseDto{
		ID:        group.ID(),
		Name:      group.Name(),
		UserIDs:   group.UserIDs(),
		EventIDs:  group.EventIDs(),
		Icon:      group.Icon(),
		CreatedAt: group.CreatedAt(),
		UpdatedAt: group.UpdatedAt(),
	}, nil
}
