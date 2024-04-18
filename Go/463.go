package main

func islandPerimeter(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				if i+1 == len(grid) || grid[i+1][j] == 0 {
					count++
				}
				if j+1 == len(grid[i]) || grid[i][j+1] == 0 {
					count++
				}
				if i-1 == -1 || grid[i-1][j] == 0 {
					count++
				}
				if j-1 == -1 || grid[i][j-1] == 0 {
					count++
				}
			}
		}
	}
	return count
}
