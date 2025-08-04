package item

import (
	"context"

	itemDomain "github.com/onion0904/CarShareSystem/app/domain/item"
)

type DeleteItemUseCase struct {
	itemRepo itemDomain.ItemRepository
}

func NewDeleteUseCase(
	itemRepo itemDomain.ItemRepository,
) *DeleteItemUseCase {
	return &DeleteItemUseCase{
		itemRepo: itemRepo,
	}
}

func (uc *DeleteItemUseCase) Run(ctx context.Context, itemID string) error {
	err := uc.itemRepo.DeleteItem(ctx, itemID)
	if err != nil {
		return err
	}
	return nil
}
