package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"

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

func (uc *CheckExistUserUseCase) Run(ctx context.Context, email string, password string) (bool, error) {
	user, err := uc.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return false, nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
