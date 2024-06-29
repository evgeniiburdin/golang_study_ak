package main

import "testing"

func Test_getVariableType(t *testing.T) {
	type args struct {
		variable interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"abcd"}, "string"},
		{"case2", args{35}, "int"},
		{"case3", args{true}, "bool"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVariableType(tt.args.variable); got != tt.want {
				t.Errorf("getVariableType() = %v, want %v", got, tt.want)
			}
		})
	}
}
