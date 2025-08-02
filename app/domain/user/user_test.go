package user

import (
	"reflect"
	"testing"
	"time"
)

func TestReconstruct(t *testing.T) {
	type args struct {
		id        string
		lastName  string
		firstName string
		email     string
		password  string
		groupIDs  []string
		eventIDs  []string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "ok case: Reconstruct & Getter",
			args: args{
				id:        "01F8B9Z6G9WBJK9XJH5M7RQK5X", // 有効なULIDの例
				lastName:  "John",
				firstName: "Doe",
				email:     "john@example.com",
				password:  "password",
			},
			want: &User{
				id:        "01F8B9Z6G9WBJK9XJH5M7RQK5X",
				lastName:  "John",
				firstName: "Doe",
				email:     "john@example.com",
				password:  "password",
				groupIDs:  nil,
				eventIDs:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reconstruct(tt.args.id, tt.args.lastName, tt.args.firstName, tt.args.email, tt.args.password, tt.args.groupIDs, tt.args.eventIDs)
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
			if got.LastName() != tt.args.lastName {
				t.Errorf("got.LastName() = %v, want %v", got.LastName(), tt.args.lastName)
			}
			if got.FirstName() != tt.args.firstName {
				t.Errorf("got.FirstName() = %v, want %v", got.FirstName(), tt.args.firstName)
			}
			if got.Email() != tt.args.email {
				t.Errorf("got.Email() = %v, want %v", got.Email(), tt.args.email)
			}
			if got.Password() != tt.args.password {
				t.Errorf("got.Password() = %v, want %v", got.Password(), tt.args.password)
			}
			if len(got.GroupIDs()) != 0 {
				t.Errorf("got.GroupIDs() = %v, want empty slice", got.GroupIDs())
			}
			if len(got.EventIDs()) != 0 {
				t.Errorf("got.EventIDs() = %v, want empty slice", got.EventIDs())
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	type args struct {
		lastName  string
		firstName string
		email     string
		password  string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "ok case: New",
			args: args{
				lastName:  "John",
				firstName: "Doe",
				email:     "john@example.com",
				password:  "password",
			},
			want: &User{
				id:        "01F8B9Z6G9WBJK9XJH5M7RQK5X",
				lastName:  "John",
				firstName: "Doe",
				email:     "john@example.com",
				password:  "password",
				groupIDs:  nil,
				eventIDs:  nil,
			},
			wantErr: false,
		},
		{
			name: "error case: small lastname",
			args: args{
				lastName:  "",
				firstName: "hogehoge",
				email:     "hogehoge",
				password:  "password",
			},
			wantErr: true,
		},
		{
			name: "error case: big lastName",
			args: args{
				lastName:  "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				firstName: "hogehoge",
				email:     "hogehoge",
				password:  "password",
			},
			wantErr: true,
		},
		{
			name: "error case: small firstname",
			args: args{
				lastName:  "hogehoge",
				firstName: "",
				email:     "hogehoge",
				password:  "password",
			},
			wantErr: true,
		},
		{
			name: "error case: big firstname",
			args: args{
				lastName:  "hogehoge",
				firstName: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				email:     "hogehoge",
				password:  "password",
			},
			wantErr: true,
		},
		{
			name: "error case: email",
			args: args{
				lastName:  "hoge",
				firstName: "hoge",
				email:     "",
				password:  "password",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.lastName, tt.args.firstName, tt.args.email, tt.args.password)
			if got != nil {
				got.id = "01F8B9Z6G9WBJK9XJH5M7RQK5X" //ulidがランダムで生成されるため
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newUser(t *testing.T) {
	type args struct {
		id        string
		lastName  string
		firstName string
		email     string
		password  string
		groupIDs  []string
		eventIDs  []string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "ok case: new",
			args: args{
				id:        "01F8B9Z6G9WBJK9XJH5M7RQK5X", // 有効なULIDの例
				lastName:  "John",
				firstName: "Doe",
				email:     "john@example.com",
				password:  "password",
			},
			want: &User{
				id:        "01F8B9Z6G9WBJK9XJH5M7RQK5X",
				lastName:  "John",
				firstName: "Doe",
				email:     "john@example.com",
				password:  "password",
				groupIDs:  nil,
				eventIDs:  nil,
			},
			wantErr: false,
		},
		{
			name: "error case: id",
			args: args{
				id:        "id",
				lastName:  "hoge",
				firstName: "hogehoge",
				email:     "hogehoge",
				password:  "password",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newUser(tt.args.id, tt.args.lastName, tt.args.firstName, tt.args.email, tt.args.password, tt.args.groupIDs, tt.args.eventIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("newUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetCreatedAt(t *testing.T) {
	user := &User{}
	now := time.Now()
	user.SetCreatedAt(now)
	if !user.CreatedAt().Equal(now) {
		t.Errorf("SetCreatedAt() failed: got %v, want %v", user.CreatedAt(), now)
	}
}

func TestUser_SetUpdatedAt(t *testing.T) {
	user := &User{}
	now := time.Now()
	user.SetUpdatedAt(now)
	if !user.UpdatedAt().Equal(now) {
		t.Errorf("SetUpdatedAt() failed: got %v, want %v", user.UpdatedAt(), now)
	}
}
