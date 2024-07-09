package main

func sortTheStudents(score [][]int, k int) [][]int {
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score)-1; j++ {
			if score[j][k] < score[j+1][k] {
				score[j], score[j+1] = score[j+1], score[j]
			}
		}
	}

	return score
}
