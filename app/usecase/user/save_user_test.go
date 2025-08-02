package user

import (
	"context"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
)

func TestSaveUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewSaveUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		dto      SaveUseCaseDto
		mockFunc func()
		want     *FindUserUseCaseDto
		wantErr  bool
	}{
		{
			name: "ok case: SaveUserUseCase",
			dto: SaveUseCaseDto{
				LastName:  "onion",
				FirstName: "gratin",
				Email:     "example@onion.com",
				Password:  "pass",
			},
			mockFunc: func() {
				mockUserRepo.EXPECT().
					Save(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, user *userDomain.User) error {
						return nil
					})

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
			//tt := tt はループ内の各テストケースが独自の tt 変数のコピーを持つようにし、
			// 並行実行時に他のテストケースの影響を受けないようにするためのイディオム
			tt := tt
			t.Parallel()
			tt.mockFunc()
			got, err := uc.Run(context.Background(), tt.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveUseCase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
