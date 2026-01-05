package main

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	smallest := math.MaxInt32
	sum := int64(0)
	neg := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < 0 {
				neg++
			}
			if int(math.Abs(float64(matrix[i][j]))) < smallest {
				smallest = int(math.Abs(float64(matrix[i][j])))
			}
			sum += int64(math.Abs(float64(matrix[i][j])))
		}
	}
	if neg%2 == 1 {
		sum -= int64(smallest) * 2
	}
	return sum
}
