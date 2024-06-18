package main

import "testing"

type sortDescIntTestCase struct {
	array [8]int
	want  [8]int
}

type sortAscIntTestCase struct {
	array [8]int
	want  [8]int
}

type sortDescFloatTestCase struct {
	array [8]float64
	want  [8]float64
}

type sortAscFloatTestCase struct {
	array [8]float64
	want  [8]float64
}

func TestSortDescInt(t *testing.T) {
	testCases := []sortDescIntTestCase{
		{[8]int{1, 6, 2, 3, 0, 4, 5, 7}, [8]int{7, 6, 5, 4, 3, 2, 1, 0}},
	}
	for _, testCase := range testCases {
		result := sortDescInt(testCase.array)
		if result != testCase.want {
			t.Errorf("sortDescInt(%v) = %v, want %v", testCase.array, result, testCase.want)
		}
	}
}

func TestSortAscInt(t *testing.T) {
	testCases := []sortAscIntTestCase{
		{[8]int{1, 6, 2, 3, 0, 4, 5, 7}, [8]int{0, 1, 2, 3, 4, 5, 6, 7}},
	}
	for _, testCase := range testCases {
		result := sortAscInt(testCase.array)
		if result != testCase.want {
			t.Errorf("sortAscInt(%v) = %v, want %v", testCase.array, result, testCase.want)
		}
	}
}

func TestSortDescFloat(t *testing.T) {
	testCases := []sortDescFloatTestCase{
		{[8]float64{1.0, 6.0, 2.0, 3.0, 0.0, 4.0, 5.0, 7.0}, [8]float64{7.0, 6.0, 5.0, 4.0, 3.0, 2.0, 1.0, 0.0}},
	}
	for _, testCase := range testCases {
		result := sortDescFloat(testCase.array)
		if result != testCase.want {
			t.Errorf("sortDescFloat(%v) = %v, want %v", testCase.array, result, testCase.want)
		}
	}
}

func TestSortAscFloat(t *testing.T) {
	testCases := []sortAscFloatTestCase{
		{[8]float64{1.0, 6.0, 2.0, 3.0, 0.0, 4.0, 5.0, 7.0}, [8]float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}},
	}
	for _, testCase := range testCases {
		result := sortAscFloat(testCase.array)
		if result != testCase.want {
			t.Errorf("sortAscFloat(%v) = %v, want %v", testCase.array, result, testCase.want)
		}
	}
}
