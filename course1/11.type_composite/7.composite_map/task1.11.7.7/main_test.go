package main

import (
	"reflect"
	"testing"
)

func Test_getUniqueUsers(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"case1", args{[]User{
			{"alice", 32, "alice@yahoo.com"},
			{"charlie", 27, "charlie@yahoo.com"},
			{"bob", 25, "bob@yahoo.com"},
			{"alice", 32, "alice@google.com"},
		}}, []User{
			{"alice", 32, "alice@yahoo.com"},
			{"charlie", 27, "charlie@yahoo.com"},
			{"bob", 25, "bob@yahoo.com"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUniqueUsers(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUniqueUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
