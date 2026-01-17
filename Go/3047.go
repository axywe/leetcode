package main

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	result := int64(0)
	for i := 0; i < len(bottomLeft); i++ {
		for j := i + 1; j < len(bottomLeft); j++ {
			x1b, y1b := bottomLeft[i][0], bottomLeft[i][1]
			x1t, y1t := topRight[i][0], topRight[i][1]
			x2b, y2b := bottomLeft[j][0], bottomLeft[j][1]
			x2t, y2t := topRight[j][0], topRight[j][1]
			minX := max(x1b, x2b)
			minY := max(y1b, y2b)
			maxX := min(x1t, x2t)
			maxY := min(y1t, y2t)
			if minX < maxX && minY < maxY {
				w := maxX - minX
				h := maxY - minY
				square := int64(min(w, h)) * int64(min(w, h))
				if square > result {
					result = square
				}
			}
		}
	}
	return result
}
