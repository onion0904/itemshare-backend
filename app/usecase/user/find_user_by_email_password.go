package user

import (
	"context"
	"time"

	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
)

type FindUserByEmailPasswordUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindUserByEmailPasswordUseCase(
	userRepo userDomain.UserRepository,
) *FindUserByEmailPasswordUseCase {
	return &FindUserByEmailPasswordUseCase{
		userRepo: userRepo,
	}
}

type FindUserByEmailPasswordUseCaseDto struct {
	ID        string
	LastName  string
	FirstName string
	Email     string
	Password  string
	Icon      string
	GroupIDs  []string
	EventIDs  []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (uc *FindUserByEmailPasswordUseCase) Run(ctx context.Context, email string) (*FindUserByEmailPasswordUseCaseDto, error) {
	user, err := uc.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &FindUserByEmailPasswordUseCaseDto{
		ID:        user.ID(),
		LastName:  user.LastName(),
		FirstName: user.FirstName(),
		Email:     user.Email(),
		Password:  user.Password(),
		Icon:      user.Icon(),
		GroupIDs:  user.GroupIDs(),
		EventIDs:  user.EventIDs(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}
