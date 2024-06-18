package main

import (
	"reflect"
	"testing"
)

func TestFilterDividers(t *testing.T) {
	type args struct {
		xs      []int
		divider int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 2}, []int{2, 4, 6, 8, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterDividers(tt.args.xs, tt.args.divider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterDividers() = %v, want %v", got, tt.want)
			}
		})
	}
}
