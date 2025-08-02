package user

import (
	"context"
	"time"

	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
)

type FindUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindUserUseCase(
	userRepo userDomain.UserRepository,
) *FindUserUseCase {
	return &FindUserUseCase{
		userRepo: userRepo,
	}
}

type FindUserUseCaseDto struct {
	ID        string
	LastName  string
	FirstName string
	Email     string
	Password  string
	GroupIDs  []string
	EventIDs  []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (uc *FindUserUseCase) Run(ctx context.Context, id string) (*FindUserUseCaseDto, error) {
	user, err := uc.userRepo.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &FindUserUseCaseDto{
		ID:        user.ID(),
		LastName:  user.LastName(),
		FirstName: user.FirstName(),
		Email:     user.Email(),
		Password:  user.Password(),
		GroupIDs:  user.GroupIDs(),
		EventIDs:  user.EventIDs(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}
