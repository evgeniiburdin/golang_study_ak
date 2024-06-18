package main

import "testing"

func Test_getType(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"emptyInterface", args{i: interface{}(nil)}, "Пустой интерфейс"},
		{"[]float64", args{i: []float64{}}, "[]float64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.args.i); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}
