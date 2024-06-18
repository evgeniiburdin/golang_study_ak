package main

import "testing"

const arraySize int = 8

type sumTestCase struct {
	array [arraySize]int
	want  int
}

type avgTestCase struct {
	array [arraySize]int
	want  float64
}

type avgFloatTestCase struct {
	array [arraySize]float64
	want  float64
}

type reverseTestCase struct {
	array [arraySize]int
	want  [arraySize]int
}

func TestSum(t *testing.T) {
	testCases := []sumTestCase{
		{[arraySize]int{1, 2, 3, 4, 5, 6, 7, 8}, 36},
		{[arraySize]int{15, 10, 5, 10, 25, 45}, 110},
	}
	for _, testCase := range testCases {
		result := sum(testCase.array)
		if result != testCase.want {
			t.Errorf("sum(%#v) = %d, want %d", testCase.array, result, testCase.want)
		}
	}
}

func TestAvg(t *testing.T) {
	testCases := []avgTestCase{
		{[arraySize]int{1, 2, 3, 4, 5, 6, 7, 8}, 4.5},
	}
	for _, testCase := range testCases {
		result := average(testCase.array)
		if result != testCase.want {
			t.Errorf("average(%#v) = %f, want %f", testCase.array, result, testCase.want)
		}
	}
}

func TestAverageFloat(t *testing.T) {
	testCases := []avgFloatTestCase{
		{[arraySize]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}, 5},
	}
	for _, testCase := range testCases {
		result := averageFloat(testCase.array)
		if result != testCase.want {
			t.Errorf("averageFloat(%#v) = %f, want %f", testCase.array, result, testCase.want)
		}
	}
}

func TestReverse(t *testing.T) {
	testCases := []reverseTestCase{
		{[arraySize]int{1, 2, 3, 4, 5, 6, 7, 8}, [arraySize]int{8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, testCase := range testCases {
		result := reverse(testCase.array)
		if result != testCase.want {
			t.Errorf("reverse(%#v) = %d, want %d", testCase.array, result, testCase.want)
		}
	}
}
