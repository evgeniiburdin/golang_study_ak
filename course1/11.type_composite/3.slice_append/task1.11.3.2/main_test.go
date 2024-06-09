package main

import (
	"reflect"
	"testing"
)

func Test_appendInt(t *testing.T) {
	type args struct {
		xs *[]int
		x  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{&[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appendInt(tt.args.xs, tt.args.x...)
			if !reflect.DeepEqual(*tt.args.xs, tt.want) {
				t.Errorf("appendInt() = %v, want %v", *tt.args.xs, tt.want)
			}
		})
	}
}
