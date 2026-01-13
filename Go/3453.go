package main

import "math"

func areaBelow(squares [][]int, cur, high, low float64) float64 {
	under := float64(0)
	for i := 0; i < len(squares); i++ {
		if float64(squares[i][1]) < cur {
			if float64(squares[i][1])+float64(squares[i][2]) < cur {
				under += float64(squares[i][2]) * float64(squares[i][2])
			} else {
				under += float64(squares[i][2]) * (cur - float64(squares[i][1]))
			}
		}
	}
	// fmt.Println("inside", under, high, low)
	return under
}

func separateSquares(squares [][]int) float64 {
	if len(squares) == 1 {
		return float64(float64(squares[0][1]) + float64(squares[0][2])/2.)
	}
	sum := 0
	minY := math.MaxInt
	maxY := math.MinInt
	for i := 0; i < len(squares); i++ {
		sum += squares[i][2] * squares[i][2]
		if squares[i][1] < minY {
			minY = squares[i][1]
		}
		if squares[i][1]+squares[i][2] > maxY {
			maxY = squares[i][1] + squares[i][2]
		}
	}

	sumFloat := float64(sum) / 2
	high := float64(maxY)
	low := float64(minY)
	cur := (high + low) / 2
	for i := 0; i < 60; i++ {
		cur = (high + low) / 2
		area := areaBelow(squares, cur, high, low)
		// fmt.Println(cur, high, low, sumFloat)
		if area >= sumFloat {
			high = cur
		} else {
			low = cur
		}
	}

	return cur
}
