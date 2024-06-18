package main

import (
	"reflect"
	"testing"
)

func Test_mergeMaps(t *testing.T) {
	type args struct {
		map1 map[string]int
		map2 map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"case1", args{map[string]int{"apple": 3, "banana": 2}, map[string]int{"orange": 5, "grape": 4}},
			map[string]int{"grape": 4, "apple": 3, "banana": 2, "orange": 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeMaps(tt.args.map1, tt.args.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
