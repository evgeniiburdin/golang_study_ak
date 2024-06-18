package main

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"emptyString", args{""}, ""},
		{"singleCharString", args{"a"}, "a"},
		{"spacesString", args{"   "}, "   "},
		{"specPuncChars", args{"Hello, world!"}, "!dlrow ,olleH"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.str); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
