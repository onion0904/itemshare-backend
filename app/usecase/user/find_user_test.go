package user

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
)

func TestFindUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewFindUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		id       string
		mockFunc func()
		want     *FindUserUseCaseDto
		wantErr  bool
	}{
		{
			name: "ok case: FindUserUseCase",
			id:   "01F8B9Z6G9WBJK9XJH5M7RQK5X",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindUser(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, id string) (*userDomain.User, error) {
						return reconstructUser(
							id,
							"onion",
							"gratin",
							"example@onion.com",
							"pass",
							nil,
							nil,
						)
					})
			},
			want: &FindUserUseCaseDto{
				ID:        "01F8B9Z6G9WBJK9XJH5M7RQK5X",
				LastName:  "onion",
				FirstName: "gratin",
				Email:     "example@onion.com",
				Password:  "pass",
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
			got, err := uc.Run(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserUseCase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func reconstructUser(
	id string,
	lastName string,
	firstName string,
	email string,
	password string,
	groupIDs []string,
	eventIDs []string,
) (*userDomain.User, error) {
	user, err := userDomain.Reconstruct(id, lastName, firstName, email, password, groupIDs, eventIDs)
	if err != nil {
		return nil, err
	}
	return user, nil
}
