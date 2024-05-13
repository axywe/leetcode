package main

func matrixScore(grid [][]int) int {
	for i := range grid {
		if grid[i][0] == 0 {
			for j := range grid[i] {
				if grid[i][j] == 0 {
					grid[i][j] = 1
				} else {
					grid[i][j] = 0
				}
			}
		}
	}
	for i := range grid[0] {
		sum := 0
		for j := range grid {
			sum += grid[j][i]
		}
		if sum*2 < len(grid) {
			for j := range grid {
				if grid[j][i] == 0 {
					grid[j][i] = 1
				} else {
					grid[j][i] = 0
				}
			}
		}
	}
	sum := 0
	for i := range grid {
		current := 0
		for _, j := range grid[i] {
			current = current << 1
			current += j
		}
		sum += current
	}
	return sum
}
