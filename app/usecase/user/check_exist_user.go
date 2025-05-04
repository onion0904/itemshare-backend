package user

import (
	"context"
	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
)

type CheckExistUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewCheckExistUserUseCase(
	userRepo userDomain.UserRepository,
) *CheckExistUserUseCase {
	return &CheckExistUserUseCase{
		userRepo: userRepo,
	}
}

func (uc *CheckExistUserUseCase) Run(ctx context.Context, email string, password string) (bool,error) {
	exist,err := uc.userRepo.ExistUser(ctx,email, password)
	if err != nil {
		// boolのゼロ値がfalseのため
        return false, err
    }
	return exist,nil
}