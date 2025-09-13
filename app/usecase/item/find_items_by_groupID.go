package item

import (
	"context"

	itemDomain "github.com/onion0904/CarShareSystem/app/domain/item"
)

type FindItemByGroupIDUseCase struct {
	itemRepo itemDomain.ItemRepository
}

func NewFindItemByGroupIDUseCase(
	itemRepo itemDomain.ItemRepository,
) *FindItemByGroupIDUseCase {
	return &FindItemByGroupIDUseCase{
		itemRepo: itemRepo,
	}
}

type FindItemByGroupIDUseCaseDto struct {
	ID      string
	Name    string
	GroupID string
}

func (uc *FindItemByGroupIDUseCase) Run(ctx context.Context, groupID string) ([]*FindItemByGroupIDUseCaseDto, error) {
	items, err := uc.itemRepo.FindItemsByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	var dto []*FindItemByGroupIDUseCaseDto
	for _, item := range *items {
		dto = append(dto,
			&FindItemByGroupIDUseCaseDto{
				ID:      item.ID(),
				Name:    item.Name(),
				GroupID: item.GroupID(),
			})
	}

	return dto, nil
}
