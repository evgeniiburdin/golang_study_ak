package main

import (
	"reflect"
	"testing"
)

func Test_appendInt(t *testing.T) {
	type args struct {
		xs []int
		x  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 2, 3}, []int{4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendInt(tt.args.xs, tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
