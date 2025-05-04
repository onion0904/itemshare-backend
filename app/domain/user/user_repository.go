package user

import "context"

type UserRepository interface {
	Update(ctx context.Context, user *User) error
    Save(ctx context.Context, user *User) error
    FindUser(ctx context.Context, UserID string) (*User, error)
    FindUserByEmailPassword(ctx context.Context, email string, password string) (*User, error)
	Delete(ctx context.Context, UserID string) error
    ExistUser(ctx context.Context,email string, password string) (bool,error)
}