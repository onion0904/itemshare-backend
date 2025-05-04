package group

import (
	"reflect"
	"testing"
	"time"
)

func TestReconstruct(t *testing.T) {
	type args struct {
		id       string
		name     string
		userIDs  []string
		eventIDs []string
		icon     string
	}
	tests := []struct {
		name    string
		args    args
		want    *Group
		wantErr bool
	}{
		{
			name: "ok case: Reconstruct & Getter",
			args: args{
				id:       "01F8B9Z6G9WBJK9XJH5M7RQK5X",  // 有効なULIDの例
				name:     "Test Group",
                icon:     "group.jpg",
			},
			want: &Group{
				id:       "01F8B9Z6G9WBJK9XJH5M7RQK5X",
                name:     "Test Group",
                userIDs:  nil,
                eventIDs: nil,
                icon:     "group.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reconstruct(tt.args.id, tt.args.name, tt.args.userIDs, tt.args.eventIDs, tt.args.icon)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reconstruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reconstruct() = %v, want %v", got, tt.want)
			}

			if got.ID() != tt.args.id {
				t.Errorf("got.ID() = %v, want %v", got.ID(), tt.args.id)
			}
			if got.Name() != tt.args.name {
				t.Errorf("got.LastName() = %v, want %v", got.Name(), tt.args.name)
			}
			if got.Icon() != tt.args.icon {
				t.Errorf("got.Icon() = %v, want %v", got.Icon(), tt.args.icon)
			}
			if len(got.UserIDs()) != 0 {
				t.Errorf("got.GroupIDs() = %v, want empty slice", got.UserIDs())
			}
			if len(got.EventIDs()) != 0 {
				t.Errorf("got.EventIDs() = %v, want empty slice", got.EventIDs())
			}
		})
	}
}

func TestNewGroup(t *testing.T) {
	type args struct {
		name    string
		userIDs []string
		icon    string
	}
	tests := []struct {
		name    string
		args    args
		want    *Group
		wantErr bool
	}{
		{
			name: "ok case: New",
			args: args{
				name:     "Test Group",
                icon:     "group.jpg",
			},
			want: &Group{
				id:       "01F8B9Z6G9WBJK9XJH5M7RQK5X",
                name:     "Test Group",
                userIDs:  nil,
                eventIDs: nil,
                icon:     "group.jpg",
			},
			wantErr: false,
		},
		{
			name: "error case: small group name",
			args: args{
				name:     "",
                icon:     "group.jpg",
			},
			wantErr: true,
		},
		{
			name: "error case: big group name",
			args: args{
				name:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                icon:     "group.jpg",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGroup(tt.args.name, tt.args.userIDs, tt.args.icon)
			if got != nil {
				got.id = "01F8B9Z6G9WBJK9XJH5M7RQK5X" //ulidがランダムで生成されるため
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newGroup(t *testing.T) {
	type args struct {
		id       string
		name     string
		userIDs  []string
		eventIDs []string
		icon     string
	}
	tests := []struct {
		name    string
		args    args
		want    *Group
		wantErr bool
	}{
		{
			name: "ok case: new",
			args: args{
				id:       "01F8B9Z6G9WBJK9XJH5M7RQK5X",   // 有効なULIDの例
				name:     "Test Group",
                icon:     "group.jpg",
			},
			want: &Group{
				id:       "01F8B9Z6G9WBJK9XJH5M7RQK5X",
                name:     "Test Group",
                userIDs:  nil,
                eventIDs: nil,
                icon:     "group.jpg",
			},
			wantErr: false,
		},
		{
			name: "error case: id",
			args: args{
				id:       "id",
				name:     "Test Group",
                icon:     "group.jpg",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newGroup(tt.args.id, tt.args.name, tt.args.userIDs, tt.args.eventIDs, tt.args.icon)
			if (err != nil) != tt.wantErr {
				t.Errorf("newGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestUser_SetCreatedAt(t *testing.T) {
	group := &Group{}
	now := time.Now()
	group.SetCreatedAt(now)
	if !group.CreatedAt().Equal(now) {
		t.Errorf("SetCreatedAt() failed: got %v, want %v", group.CreatedAt(), now)
	}
}

func TestUser_SetUpdatedAt(t *testing.T) {
	group := &Group{}
	now := time.Now()
	group.SetUpdatedAt(now)
	if !group.UpdatedAt().Equal(now) {
		t.Errorf("SetUpdatedAt() failed: got %v, want %v", group.UpdatedAt(), now)
	}
}