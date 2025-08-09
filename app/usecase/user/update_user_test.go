// これは未実装です

package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/samber/lo"
	"go.uber.org/mock/gomock"

	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
)

func TestUpdateUseCaseRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewUpdateUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		id       string
		dto      UpdateUseCaseDto
		mockFunc func()
		want     *FindUserUseCaseDto
		wantErr  bool
	}{
		{
			name: "ok case: UpdateUserUseCase",
			id:   "01F8B9Z6G9WBJK9XJH5M7RQK5X",
			dto: UpdateUseCaseDto{
				LastName:  lo.ToPtr("Updated"),
				FirstName: lo.ToPtr("User"),
			},
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindUser(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, id string) (*userDomain.User, error) {
						return reconstructUser(
							"01F8B9Z6G9WBJK9XJH5M7RQK5X",
							"onion",
							"gratin",
							"example@onion.com",
							"pass",
							nil,
							nil,
						)
					})

				mockUserRepo.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, user *userDomain.User) error {
						return nil
					})

				mockUserRepo.
					EXPECT().
					FindUser(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, id string) (*userDomain.User, error) {
						return reconstructUser(
							"01F8B9Z6G9WBJK9XJH5M7RQK5X",
							"Updated",
							"User",
							"updated@example.com",
							"pass",
							nil,
							nil,
						)
					})
			},
			want: &FindUserUseCaseDto{
				ID:        "01F8B9Z6G9WBJK9XJH5M7RQK5X",
				LastName:  "Updated",
				FirstName: "User",
				Email:     "updated@example.com",
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
			got, err := uc.Run(context.Background(), tt.id, tt.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateUseCase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
