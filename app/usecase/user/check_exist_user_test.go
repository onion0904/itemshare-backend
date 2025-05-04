package user

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
)

func TestCheckExistUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewCheckExistUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		email    string
		password string
		mockFunc func()
		want    bool
		wantErr bool
	}{
		{
			name: "ok case: FindUserUseCase",
			email: "example@onion.com",
			password: "pass",
			mockFunc: func()  {
				mockUserRepo.
					EXPECT().
					ExistUser(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, email string, password string) (bool, error){
						return false,nil
					})
			},
			want: false,
			wantErr: false,
		},
		{
			name: "error case: FindUserUseCase",
			email: "example@onion.com",
			password: "pass",
			mockFunc: func()  {
				mockUserRepo.
					EXPECT().
					ExistUser(gomock.Any(), gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, email string, password string) (bool, error){
						return false,errDomain.NewError("何らかの原因")
					})
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.Run(context.Background(), tt.email, tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckExistUserUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckExistUserUseCase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
