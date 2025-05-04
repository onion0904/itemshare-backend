package group

import (
	"unicode/utf8"
	"github.com/onion0904/CarShareSystem/pkg/ulid"
	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
	"time"
)


type Group struct {
	id string
	name string
	userIDs []string
	eventIDs []string
	icon string
	createdAt time.Time
	updatedAt time.Time
}


func Reconstruct(
	id string,
	name string,
	userIDs []string,
	eventIDs []string,
	icon string,
) (*Group, error) {
	return newGroup(
		id,
		name,
		userIDs,
		eventIDs,
		icon,
	)
}

func NewGroup(
	name string,
	userIDs []string,
	icon string,
) (*Group, error) {
	return newGroup(
		ulid.NewUlid(),
		name,
		userIDs,
		nil,
		icon,
	)
}

func newGroup(
	id string,
	name string,
	userIDs []string,
	eventIDs []string,
	icon string,
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
		id:          id,
		name:        name,
		userIDs:     userIDs,
		eventIDs:    eventIDs,
		icon:        icon,
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

func (p *Group) Icon() string {
    return p.icon
}

func (p *Group) CreatedAt() time.Time {
    return p.createdAt
}

func (p *Group) UpdatedAt() time.Time {
    return p.updatedAt
}

func (u *Group) SetCreatedAt(t time.Time){
	u.createdAt = t
}

func (u *Group) SetUpdatedAt(t time.Time){
    u.updatedAt = t
}


const (
	nameLengthMin = 1
	nameLengthMax = 100
)