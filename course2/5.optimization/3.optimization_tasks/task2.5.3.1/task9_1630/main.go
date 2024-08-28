package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	answer := make([]bool, 0, len(l))
	var isArithmetic bool
	var numsCopy []int
	for i := 0; i < len(l); i++ {
		numsCopy = make([]int, r[i]-l[i]+1)
		copy(numsCopy, nums[l[i]:r[i]+1])
		sort.Ints(numsCopy)
		isArithmetic = true
		for j := 0; j < len(numsCopy)-2; j++ {
			if numsCopy[j]-numsCopy[j+1] != numsCopy[j+1]-numsCopy[j+2] {
				isArithmetic = false
				break
			}
		}
		if !isArithmetic {
			answer = append(answer, false)
			continue
		}
		answer = append(answer, true)
	}
	return answer
}

func main() {
	fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
}
