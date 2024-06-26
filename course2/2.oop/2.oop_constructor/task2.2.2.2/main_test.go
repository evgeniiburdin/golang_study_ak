package main

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		ID      int
		options []UserOption
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{name: "test1",
			args: args{
				ID:      3,
				options: []UserOption{WithUsername("testuser"), WithEmail("testuser@domain.com"), WithRole("admin")},
			},
			want: &User{
				ID:       3,
				Username: "testuser",
				Email:    "testuser@domain.com",
				Role:     "admin"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.ID, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithEmail(t *testing.T) {
	type args struct {
		Email string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{name: "test1",
			args: args{
				Email: "testuser@domain.com",
			},
			want: User{ID: 3, Email: "testuser@domain.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := WithEmail(tt.args.Email)
			user := *NewUser(3, option)
			if got := user; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRole(t *testing.T) {
	type args struct {
		Role string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "test1",
			args: args{
				Role: "admin",
			},
			want: User{ID: 3, Role: "admin"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := WithRole(tt.args.Role)
			user := *NewUser(3, option)
			if got := user; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithUsername(t *testing.T) {
	type args struct {
		Username string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "test1",
			args: args{
				Username: "testuser",
			},
			want: User{ID: 3, Username: "testuser"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := WithUsername(tt.args.Username)
			user := *NewUser(3, option)
			if got := user; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
