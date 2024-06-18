package main

import "testing"

func Test_countBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{"Привет, мир!"}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBytes(tt.args.s); got != tt.want {
				t.Errorf("countBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSymbols(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{"Привет, мир!"}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSymbols(tt.args.s); got != tt.want {
				t.Errorf("countSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}
