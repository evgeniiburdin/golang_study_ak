package main

import (
	"reflect"
	"testing"
)

func TestInsertAfterIDX(t *testing.T) {
	type args struct {
		xs  []int
		idx int
		x   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1",
			args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 3, []int{0, 0, 0, 0, 0}},
			[]int{1, 2, 3, 4, 0, 0, 0, 0, 0, 5, 6, 7, 8, 9, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertAfterIDX(tt.args.xs, tt.args.idx, tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAfterIDX() = %v, want %v", got, tt.want)
			}
		})
	}
}
