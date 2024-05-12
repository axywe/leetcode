package main

func largestLocal(grid [][]int) [][]int {
	n := len(grid) - 2
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			max := 0
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if grid[i+k][j+l] > max {
						max = grid[i+k][j+l]
					}
				}
			}
			result[i][j] = max
		}
	}
	return result
}
