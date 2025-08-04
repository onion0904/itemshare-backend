package item

import (
	"context"

	itemDomain "github.com/onion0904/CarShareSystem/app/domain/item"
)

type FindItemByIDUseCase struct {
	itemRepo itemDomain.ItemRepository
}

func NewFindItemByIDUseCase(
	itemRepo itemDomain.ItemRepository,
) *FindItemByIDUseCase {
	return &FindItemByIDUseCase{
		itemRepo: itemRepo,
	}
}

type FindItemByIDUseCaseDto struct {
	ID          string
	Name      string
	GroupID      string
}

func (uc *FindItemByIDUseCase) Run(ctx context.Context, itemID string) (*FindItemByIDUseCaseDto, error) {
	item, err := uc.itemRepo.FindItemByID(ctx, itemID)
	if err != nil {
		return nil, err
	}
	return &FindItemByIDUseCaseDto{
		ID:          item.ID(),
		Name: 		 item.Name(),
		GroupID:     item.GroupID(),
	}, nil
}
