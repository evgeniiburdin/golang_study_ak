package main

import (
	"reflect"
	"testing"
)

func Test_getStringHeader(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{"Hello, World!"}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStringHeader(tt.args.s); !reflect.DeepEqual(got.Len, tt.want) {
				t.Errorf("getStringHeader().Len = %v, want %v", got.Len, tt.want)
			}
		})
	}
}
