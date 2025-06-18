package user

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
)

func TestFindUserByEmailUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewFindUserByEmailPasswordUseCase(mockUserRepo)

	tests := []struct {
		name     string
		email    string
		mockFunc func()
		want     *FindUserByEmailPasswordUseCaseDto
		wantErr  bool
	}{
		{
			name:  "ok case: FindUserUseCase",
			email: "example@onion.com",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindUserByEmail(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, email string) (*userDomain.User, error) {
						return reconstructUser(
							"01F8B9Z6G9WBJK9XJH5M7RQK5X",
							"onion",
							"gratin",
							email,
							"password",
							"icon",
							nil,
							nil,
						)
					})
			},
			want: &FindUserByEmailPasswordUseCaseDto{
				ID:        "01F8B9Z6G9WBJK9XJH5M7RQK5X",
				LastName:  "onion",
				FirstName: "gratin",
				Email:     "example@onion.com",
				Password:  "password",
				Icon:      "icon",
				GroupIDs:  nil,
				EventIDs:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.Run(context.Background(), tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByEmailPasswordUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserByEmailPasswordUseCase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
