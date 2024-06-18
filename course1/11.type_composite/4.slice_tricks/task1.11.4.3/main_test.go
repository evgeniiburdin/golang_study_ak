package main

import (
	"testing"
)

func TestRemoveExtraMemory(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name         string
		args         args
		wantSliceCap int
	}{
		{"case1", args{make([]int, 3, 12)}, 3},
		{"case2", args{make([]int, 16, 32)}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveExtraMemory(tt.args.xs); cap(got) != tt.wantSliceCap {
				t.Errorf("RemoveExtraMemory() = %v, cap = %v, want %v", got, cap(got), tt.wantSliceCap)
			}
		})
	}
}
