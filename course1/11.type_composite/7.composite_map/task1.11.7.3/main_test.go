package main

import "testing"

func Test_createUniqueText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"bar bar bar foo foo baz"}, "bar foo baz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createUniqueText(tt.args.text); got != tt.want {
				t.Errorf("createUniqueText() = %v, want %v", got, tt.want)
			}
		})
	}
}
