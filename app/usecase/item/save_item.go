package item

import (
	"context"

	itemDomain "github.com/onion0904/CarShareSystem/app/domain/item"
)

type SaveUseCase struct {
	itemRepo itemDomain.ItemRepository
}

func NewSaveItemUseCase(
	itemRepo itemDomain.ItemRepository,
) *SaveUseCase {
	return &SaveUseCase{
		itemRepo: itemRepo,
	}
}

type SaveUseCaseDto struct {
	Name    string
	GroupID string
}

func (uc *SaveUseCase) Run(ctx context.Context, dto SaveUseCaseDto) (*FindItemByIDUseCaseDto, error) {
	// dtoからitemへ変換
	nitem, err := itemDomain.NewItem(dto.Name, dto.GroupID)
	if err != nil {
		return nil, err
	}
	err = uc.itemRepo.SaveItem(ctx, nitem)
	if err != nil {
		return nil, err
	}
	item, err := uc.itemRepo.FindItemByID(ctx, nitem.ID())
	if err != nil {
		return nil, err
	}
	return &FindItemByIDUseCaseDto{
		ID:      item.ID(),
		Name:    item.Name(),
		GroupID: item.GroupID(),
	}, nil
}
