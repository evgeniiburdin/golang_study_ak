package main

import (
	"reflect"
	"testing"
)

func TestInsertToStart(t *testing.T) {
	type args struct {
		xs []int
		x  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5}, nil}, []int{1, 2, 3, 4, 5}},
		{"case2", args{[]int{1, 2, 3, 4, 5}, []int{0, 0, 0, 0, 0}},
			[]int{0, 0, 0, 0, 0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertToStart(tt.args.xs, tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertToStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
