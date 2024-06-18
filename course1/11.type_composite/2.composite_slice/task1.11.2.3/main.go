package main

func bitwiseXOR(n, res int) int {
	return n ^ res
}

func findSingleNumber(numbers []int) int {
	res := 0
	for _, num := range numbers {
		res = bitwiseXOR(num, res)
	}
	return res
}
