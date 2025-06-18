package user

import "context"

type UserRepository interface {
	Update(ctx context.Context, user *User) error
	Save(ctx context.Context, user *User) error
	FindUser(ctx context.Context, UserID string) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, UserID string) error
}
