package user

import (
	"net/mail"
	"time"
	"unicode/utf8"

	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
	"github.com/onion0904/CarShareSystem/pkg/ulid"
)

type User struct {
	id        string
	lastName  string
	firstName string
	email     string
	password  string
	icon      string
	groupIDs  []string
	eventIDs  []string
	createdAt time.Time
	updatedAt time.Time
}

func Reconstruct(
	id string,
	lastName string,
	firstName string,
	email string,
	password string,
	icon string,
	groupIDs []string,
	eventIDs []string,
) (*User, error) {
	return newUser(
		id,
		lastName,
		firstName,
		email,
		password,
		icon,
		groupIDs,
		eventIDs,
	)
}

func NewUser(
	lastName string,
	firstName string,
	email string,
	password string,
	icon string,
) (*User, error) {
	return newUser(
		ulid.NewUlid(),
		lastName,
		firstName,
		email,
		password,
		icon,
		nil,
		nil,
	)
}

func newUser(
	id string,
	lastName string,
	firstName string,
	email string,
	password string,
	icon string,
	groupIDs []string,
	eventIDs []string,
) (*User, error) {
	// ownerIDのバリデーション
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("オーナーIDの値が不正です。")
	}
	// 名前のバリデーション
	if utf8.RuneCountInString(lastName) < nameLengthMin || utf8.RuneCountInString(lastName) > nameLengthMax {
		return nil, errDomain.NewError("名前（姓）の値が不正です。")
	}
	if utf8.RuneCountInString(firstName) < nameLengthMin || utf8.RuneCountInString(firstName) > nameLengthMax {
		return nil, errDomain.NewError("名前（名）の値が不正です。")
	}

	// メールアドレスのバリデーション
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errDomain.NewError("メールアドレスの値が不正です。")
	}

	return &User{
		id:        id,
		lastName:  lastName,
		firstName: firstName,
		email:     email,
		password:  password,
		icon:      icon,
		groupIDs:  groupIDs,
		eventIDs:  eventIDs,
	}, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Icon() string {
	return u.icon
}

func (u *User) GroupIDs() []string {
	return u.groupIDs
}

func (u *User) EventIDs() []string {
	return u.eventIDs
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) SetCreatedAt(t time.Time) {
	u.createdAt = t
}

func (u *User) SetUpdatedAt(t time.Time) {
	u.updatedAt = t
}

const (
	nameLengthMin = 1
	nameLengthMax = 50
)
