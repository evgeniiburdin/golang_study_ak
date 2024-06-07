package main

import (
	"math/rand"
	"reflect"

	"testing"
)

type testCase struct {
	slice    []float64
	expected float64
}

func generateSlice(size int) []float64 {
	slice := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Float64())
	}
	return slice
}

func TestAverage(t *testing.T) {
	sliceLen := 15
	reflect.DeepEqual(generateSlice(sliceLen), generateSlice(sliceLen))

	testCases := []testCase{
		{[]float64{1.1, 2.2, 3.3, 4.4}, 2.75},
		{[]float64{1.8, 2.4, 3.2, 4.4}, 2.95},
		{[]float64{7.0, 8.2, 15.3, 1.4}, 7.975},
	}

	for _, tc := range testCases {
		result := average(tc.slice)
		if result != tc.expected {
			t.Errorf("average(%v) = %f, want %f", tc.slice, result, tc.expected)
		}
	}
}
