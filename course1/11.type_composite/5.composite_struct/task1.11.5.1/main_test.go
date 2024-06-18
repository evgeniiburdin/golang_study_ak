package main

import (
	"testing"
)

func Test_getUsers(t *testing.T) {
	if got := getUsers(); len(got) != 10 {
		t.Errorf("getUsers() = %v, want 10 elements", got)
	}
}

func Test_preparePrint(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{[]User{{"Alice", 31}, {"Bob", 54}}},
			"Имя: Alice, Возраст: 31\nИмя: Bob, Возраст: 54\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preparePrint(tt.args.users); got != tt.want {
				t.Errorf("preparePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
