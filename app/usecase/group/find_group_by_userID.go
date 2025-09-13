package group

import (
	"context"
	"time"

	groupDomain "github.com/onion0904/CarShareSystem/app/domain/group"
)

type FindGroupsByUserIDUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewFindGroupsByUserIDUseCase(
	groupRepo groupDomain.GroupRepository,
) *FindGroupsByUserIDUseCase {
	return &FindGroupsByUserIDUseCase{
		groupRepo: groupRepo,
	}
}

type FindGroupsByUserIDUseCaseDto struct {
	ID        string
	Name      string
	UserIDs   []string
	EventIDs  []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (uc *FindGroupsByUserIDUseCase) Run(ctx context.Context, userID string) ([]*FindGroupsByUserIDUseCaseDto, error) {
	groups, err := uc.groupRepo.FindGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]*FindGroupsByUserIDUseCaseDto, 0, len(groups))
	for _, group := range groups {
		if group != nil {
			result = append(result, &FindGroupsByUserIDUseCaseDto{
				ID:        group.ID(),
				Name:      group.Name(),
				UserIDs:   group.UserIDs(),
				EventIDs:  group.EventIDs(),
				CreatedAt: group.CreatedAt(),
				UpdatedAt: group.UpdatedAt(),
			})
		}
	}

	return result, nil
}
