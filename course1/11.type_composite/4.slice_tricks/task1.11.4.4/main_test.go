package main

import (
	"reflect"
	"testing"
)

func TestRemoveIDX(t *testing.T) {
	type args struct {
		xs  []int
		idx int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5}, 3}, []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveIDX(tt.args.xs, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIDX() = %v, want %v", got, tt.want)
			}
		})
	}
}
