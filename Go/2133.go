package main

func checkValid(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		m1 := make(map[int]bool)
		m2 := make(map[int]bool)
		for j := 0; j < len(matrix[i]); j++ {
			if m1[matrix[i][j]] {
				return false
			} else {
				m1[matrix[i][j]] = true
			}
			if m2[matrix[j][i]] {
				return false
			} else {
				m2[matrix[j][i]] = true
			}
		}
	}
	return true
}
