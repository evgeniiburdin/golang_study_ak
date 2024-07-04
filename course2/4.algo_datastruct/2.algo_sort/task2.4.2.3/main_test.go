package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		arr1 []User
		arr2 []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{
			name: "mergeTest",
			args: args{
				arr1: []User{{Name: "Alice", Age: 25}, {Name: "Charlie", Age: 35}},
				arr2: []User{{Name: "Bob", Age: 30}, {Name: "Dave", Age: 40}},
			},
			want: []User{{Name: "Alice", Age: 25}, {Name: "Bob", Age: 30}, {Name: "Charlie", Age: 35}, {Name: "Dave", Age: 40}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
