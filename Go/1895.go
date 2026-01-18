package main

func colSum(grid [][]int, x, y, size int) int {
	sum := -1
	for j := y; j <= y+size; j++ {
		colSum := 0
		for i := x; i <= x+size; i++ {
			colSum += grid[i][j]
		}
		if sum == -1 {
			sum = colSum
		} else if sum != -1 && sum != colSum {
			sum = -2
		}
	}
	return sum
}

func rowSum(grid [][]int, x, y, size int) int {
	sum := -1
	for i := x; i <= x+size; i++ {
		rowSum := 0
		for j := y; j <= y+size; j++ {
			rowSum += grid[i][j]
		}
		if sum == -1 {
			sum = rowSum
		} else if sum != -1 && sum != rowSum {
			sum = -2
		}
	}
	return sum
}

func diagSum(grid [][]int, x, y, size int) int {
	sumMain := 0
	sumSec := 0
	for i := 0; i <= size; i++ {
		sumMain += grid[x+i][y+i]
		sumSec += grid[x+i][y+size-i]
	}
	if sumMain == sumSec {
		return sumMain
	}
	return -2
}

func findMagicSquare(grid [][]int, x, y int) int {
	min := min(len(grid)-x, len(grid[0])-y)
	maxSize := 1
	for i := 1; i < min; i++ {
		r := rowSum(grid, x, y, i)
		c := colSum(grid, x, y, i)
		d := diagSum(grid, x, y, i)
		if r != -2 && c != -2 && d != -2 && r == c && r == d {
			maxSize = i + 1
		}
	}
	return maxSize
}

func largestMagicSquare(grid [][]int) int {
	result := 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			size := findMagicSquare(grid, i, j)
			result = max(result, size)
		}
	}
	return result
}
