package user

import (
	"context"
	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
)

type DeleteUseCase struct {
	userRepo userDomain.UserRepository
}

func NewDeleteUseCase(
	userRepo userDomain.UserRepository,
) *DeleteUseCase {
	return &DeleteUseCase{
		userRepo: userRepo,
	}
}

func (uc *DeleteUseCase) Run(ctx context.Context, id string) error {
	err := uc.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
