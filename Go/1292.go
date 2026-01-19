package main

func matrixSum(mat [][]int, x, y, size int) int {
	sum := 0
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			sum += mat[i][j]
		}
	}
	return sum
}

func maxSideLength(mat [][]int, threshold int) int {
	curSize := 1
	for i := 0; i <= len(mat)-curSize; i++ {
		for j := 0; j <= len(mat[i])-curSize && i <= len(mat)-curSize; j++ {
			sum := matrixSum(mat, i, j, curSize)
			if sum <= threshold {
				curSize++
				j--
			}
		}
	}
	return curSize - 1
}
