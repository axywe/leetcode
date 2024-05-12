package main

func satisfiesConditions(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != grid[0][j] {
				return false
			}
			if j+1 < len(grid[i]) && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}
