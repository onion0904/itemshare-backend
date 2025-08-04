package item

import "github.com/onion0904/CarShareSystem/pkg/ulid"

type Item struct {
	id         string
	name 	   string
	groupID    string
}

func Reconstruct(
	id         string,
	name 	   string,
	groupID    string,
) (*Item, error) {
	return newItem(
		id,
		name,
		groupID,
	)
}

func NewItem(
	name 	   string,
	groupID    string,
) (*Item, error) {
	return newItem(
		ulid.NewUlid(),
		name,
		groupID,
	)
}

func newItem(
	id         string,
	name 	   string,
	groupID    string,
) (*Item, error) {
	return &Item{
		id: 		id,
		name: 		name,
		groupID: 	groupID,
	}, nil
}

func (i *Item) ID() string {
	return i.id
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) GroupID() string {
	return i.groupID
}