package repository

import (
	"context"
	"database/sql"
	"errors"

	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
	"github.com/onion0904/CarShareSystem/app/domain/item"
	"github.com/onion0904/CarShareSystem/app/infrastructure/db"
	dbgen "github.com/onion0904/CarShareSystem/app/infrastructure/db/sqlc/dbgen"
)

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) item.ItemRepository {
	return &itemRepository{db: db}
}

func (ir *itemRepository) SaveItem(ctx context.Context, item *item.Item) error {
	query := db.GetQuery(ctx)

	err := query.InsertItem(ctx, dbgen.InsertItemParams{
		ID:      item.ID(),
		GroupID: item.GroupID(),
		Name:    item.Name(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ir *itemRepository) DeleteItem(ctx context.Context, itemID string) error {
	query := db.GetQuery(ctx)

	err := query.DeleteItem(ctx, itemID)
	if err != nil {
		return err
	}
	return nil
}

func (ir *itemRepository) FindItemByID(ctx context.Context, itemID string) (*item.Item, error) {
	query := db.GetQuery(ctx)

	i, err := query.GetItemByID(ctx, itemID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NewError("Item not found")
		}
		return nil, err
	}

	ni, err := item.Reconstruct(
		i.ID,
		i.Name,
		i.GroupID,
	)
	if err != nil {
		return nil, err
	}
	return ni, nil
}

func (ir *itemRepository) FindItemsByGroupID(ctx context.Context, groupID string) (*[]item.Item, error) {
	query := db.GetQuery(ctx)

	is, err := query.GetItemsByGroupID(ctx, groupID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, nil
	}

	var result []item.Item
	for _, i := range is {
		ni, err := item.Reconstruct(
			i.ID,
			i.Name,
			i.GroupID,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, *ni)
	}
	return &result, nil
}
