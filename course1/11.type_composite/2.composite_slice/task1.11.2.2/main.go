package main

func MaxDifference(numbers []int) int {
	min, max := numbers[0], numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max - min
}
