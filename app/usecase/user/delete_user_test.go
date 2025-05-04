package user

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"

	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
	errDomain "github.com/onion0904/CarShareSystem/app/domain/error"
)

func TestDeleteUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewDeleteUseCase(mockUserRepo)

	tests := []struct {
		name    string
		id       string
		mockFunc func()
		wantErr bool
	}{
		{
			name: "ok case: DeleteUser",
			id: "01F8B9Z6G9WBJK9XJH5M7RQK5X",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					DoAndReturn(func (ctx context.Context, id string) error {
						return nil
					})
			},
			wantErr: false,
		},
		{
			name: "error case: DeleteUser",
			id: "hoge",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					Delete(gomock.Any(), gomock.Any()).
					DoAndReturn(func (ctx context.Context, id string) error {
						return errDomain.NewError("ユーザーが存在しませんとか、")
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
			if err := uc.Run(context.Background(), tt.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
