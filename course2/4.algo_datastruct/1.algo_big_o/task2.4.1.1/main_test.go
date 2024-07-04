package main

import (
	"reflect"
	"testing"
)

func Test_compareWhichFactorialIsFaster(t *testing.T) {
	tests := []struct {
		name string
		want map[string]bool
	}{
		{
			name: "codeMustBe100%CoveredWithTests",
			want: map[string]bool{
				"recursive": false,
				"iterative": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareWhichFactorialIsFaster(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareWhichFactorialIsFaster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorialIterative(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testFactorial5",
			args: args{n: 5},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialIterative(tt.args.n); got != tt.want {
				t.Errorf("factorialIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorialRecursive(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testFactorial5",
			args: args{n: 5},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialRecursive(tt.args.n); got != tt.want {
				t.Errorf("factorialRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
