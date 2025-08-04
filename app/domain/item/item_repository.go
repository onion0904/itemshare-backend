package item

import "context"

type ItemRepository interface {
	SaveItem(ctx context.Context, item *Item) error
	FindItemByID(ctx context.Context, itemID string) (*Item, error)
	FindItemsByGroupID(ctx context.Context, groupID string) (*[]Item, error)
	DeleteItem(ctx context.Context, itemID string) error
}