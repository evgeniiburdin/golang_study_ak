package main

import (
	"reflect"
	"testing"
)

func TestGetSubSlice(t *testing.T) {
	type testCase = struct {
		slice []int
		start int
		end   int
		want  []int
	}
	testCases := []testCase{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 3, 7, []int{4, 5, 6, 7}},
	}
	for _, tc := range testCases {
		result := getSubSlice(tc.slice, tc.start, tc.end)
		if reflect.DeepEqual(result, tc.want) == false {
			t.Errorf("getSubSlice(%v, %v, %v) got %v want %v", tc.slice, tc.start, tc.end, result, tc.want)
		}
	}
}
