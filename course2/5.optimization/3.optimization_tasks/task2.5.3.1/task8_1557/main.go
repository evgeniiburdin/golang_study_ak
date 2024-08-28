package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	isReachable := make(map[int]bool, n)
	for _, edge := range edges {
		isReachable[edge[1]] = true
		if _, ok := isReachable[edge[0]]; !ok {
			isReachable[edge[0]] = false
		}
	}

	result := make([]int, 0)
	for vertex, reachable := range isReachable {
		if !reachable {
			result = append(result, vertex)
		}
	}

	return result
}
