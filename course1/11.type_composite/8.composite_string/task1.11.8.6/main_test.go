package main

import "testing"

func TestCountVowels(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{"hello world"}, 3},
		{"case2", args{"aeiou"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountVowels(tt.args.str); got != tt.want {
				t.Errorf("CountVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
