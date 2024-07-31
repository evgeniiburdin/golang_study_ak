package main

import (
	"reflect"
	"testing"
)

func Test_countRussianLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{"case1", args{"Привет, мир!"}, map[rune]int{'в': 1, 'е': 1, 'т': 1, 'м': 1, 'П': 1, 'р': 2, 'и': 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRussianLetters(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countRussianLetters() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}

func Test_isRussianLetter(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{rune('ы')}, true},
		{"case2", args{rune('Ж')}, true},
		{"case3", args{rune('G')}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRussianLetter(tt.args.r); got != tt.want {
				t.Errorf("isRussianLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
