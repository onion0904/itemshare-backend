package group

import (
	"time"
	"unicode/utf8"

	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
	"github.com/onion0904/CarShareSystem/pkg/ulid"
)

type Group struct {
	id        string
	name      string
	userIDs   []string
	eventIDs  []string
	createdAt time.Time
	updatedAt time.Time
}

func Reconstruct(
	id string,
	name string,
	userIDs []string,
	eventIDs []string,
) (*Group, error) {
	return newGroup(
		id,
		name,
		userIDs,
		eventIDs,
	)
}

func NewGroup(
	name string,
	userIDs []string,
) (*Group, error) {
	return newGroup(
		ulid.NewUlid(),
		name,
		userIDs,
		nil,
	)
}

func newGroup(
	id string,
	name string,
	userIDs []string,
	eventIDs []string,
) (*Group, error) {
	// ownerIDのバリデーション
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("オーナーIDの値が不正です。")
	}
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("グループ名が不正です。")
	}

	return &Group{
		id:       id,
		name:     name,
		userIDs:  userIDs,
		eventIDs: eventIDs,
	}, nil
}

func (p *Group) ID() string {
	return p.id
}

func (p *Group) Name() string {
	return p.name
}

func (p *Group) UserIDs() []string {
	return p.userIDs
}

func (p *Group) EventIDs() []string {
	return p.eventIDs
}

func (p *Group) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Group) UpdatedAt() time.Time {
	return p.updatedAt
}

func (u *Group) SetCreatedAt(t time.Time) {
	u.createdAt = t
}

func (u *Group) SetUpdatedAt(t time.Time) {
	u.updatedAt = t
}

const (
	nameLengthMin = 1
	nameLengthMax = 100
)
