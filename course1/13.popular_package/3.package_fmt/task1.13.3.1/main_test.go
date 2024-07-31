package main

import "testing"

func Test_generateMathString(t *testing.T) {
	type args struct {
		operands []int
		operator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{[]int{2, 4, 6}, "-"}, "2 - 4 - 6 = -8"},
		{"case2", args{[]int{2, 4, 6}, "+"}, "2 + 4 + 6 = 12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMathString(tt.args.operands, tt.args.operator); got != tt.want {
				t.Errorf("generateMathString() = %v, want %v", got, tt.want)
			}
		})
	}
}
