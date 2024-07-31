package main

import (
	"fmt"
	"time"
)

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	var result int = 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}

func compareWhichFactorialIsFaster() map[string]bool {
	var n int = 10000000

	startTime := time.Now()
	factorialRecursive(n)
	recursiveTime := time.Since(startTime)

	startTime = time.Now()
	factorialIterative(n)
	iterativeTime := time.Since(startTime)

	fmt.Println("recursive: ", recursiveTime.Seconds())
	fmt.Println("iterative: ", iterativeTime.Seconds())

	resultMap := make(map[string]bool)
	if recursiveTime > iterativeTime {
		resultMap["recursive"] = false
		resultMap["iterative"] = true
		return resultMap
	}
	if recursiveTime < iterativeTime {
		resultMap["recursive"] = true
		resultMap["iterative"] = false
		return resultMap
	}
	return map[string]bool{"recursive": false, "iterative": false}
}

/*
func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("OS/Arch:", runtime.GOOS, "/", runtime.GOARCH)

	fmt.Println("Which factorial is faster?")
	fmt.Println(compareWhichFactorialIsFaster())
}
*/
