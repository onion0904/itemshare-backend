package user

import (
	"context"
	"reflect"
	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
)

type UpdateUseCase struct {
	userRepo userDomain.UserRepository
}

func NewUpdateUserUseCase(
	userRepo userDomain.UserRepository,
) *UpdateUseCase {
	return &UpdateUseCase{
		userRepo: userRepo,
	}
}

type UpdateUseCaseDto struct {
	LastName *string
	FirstName *string
	Email *string
	Icon *string
}

func (uc *UpdateUseCase) Run(ctx context.Context, id string, dto UpdateUseCaseDto) (*FindUserUseCaseDto,error) {
	user ,err := uc.userRepo.FindUser(ctx,id)
	if err != nil {
        return nil,err
    }
	defaultValue(&dto)
	
	nuser, err := userDomain.Reconstruct(id,*dto.LastName, *dto.FirstName, *dto.Email, user.Password(), *dto.Icon,user.GroupIDs(),user.EventIDs())
	if err != nil {
		return nil,err
	}
	err = uc.userRepo.Update(ctx, nuser)
	if err != nil {
		return nil, err
	}
	updatedUser, err := uc.userRepo.FindUser(ctx, user.ID())
	if err != nil {
		return nil, err
	}
	return &FindUserUseCaseDto{
		ID:          updatedUser.ID(),
		LastName:    updatedUser.LastName(),
		FirstName:   updatedUser.FirstName(),
		Email:       updatedUser.Email(),
		Password:    updatedUser.Password(),
		Icon:        updatedUser.Icon(),
		GroupIDs:    updatedUser.GroupIDs(),
		EventIDs:    updatedUser.EventIDs(),
		CreatedAt:   updatedUser.CreatedAt(),
        UpdatedAt:   updatedUser.UpdatedAt(),
	}, nil
}

// inputの構造体のnilの部分をデフォルト値に変換
func defaultValue (inDTO *UpdateUseCaseDto){
	v := reflect.ValueOf(&inDTO).Elem()
	for i:=0; i<v.NumField(); i++{
		field := v.Field(i)
		if field.IsNil(){
			defaultValue := ""
			field.Set(reflect.ValueOf(&defaultValue))
		}
	}
}