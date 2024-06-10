package main

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	type args struct {
		xs []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"case1", args{[]interface{}{"a", "b", "c"}}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.xs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperate(t *testing.T) {
	type args struct {
		f func(xs ...interface{}) interface{}
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"caseConcat",
			args{
				Concat,
				[]interface{}{"Hello, ", "World!"},
			},
			"Hello, World!",
		},
		{
			"caseSumInts",
			args{
				Sum,
				[]interface{}{1, 2, 3, 4, 5},
			},
			15,
		},
		{
			"caseSumFloats",
			args{
				Sum,
				[]interface{}{1.1, 2.2, 3.3, 4.4, 5.5},
			},
			16.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Operate(tt.args.f, tt.args.i...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		xs []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"caseInts", args{[]interface{}{1, 2, 3, 4, 5}}, 15},
		{"caseFloats", args{[]interface{}{1.1, 2.2, 3.3, 4.4, 5.5}}, 16.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.xs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
