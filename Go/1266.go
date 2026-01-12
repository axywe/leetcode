package main

func minTimeToVisitAllPoints(points [][]int) int {
	seconds := 0
	for i := 0; i < len(points)-1; i++ {
		dx := points[i][0] - points[i+1][0]
		dy := points[i][1] - points[i+1][1]
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		seconds += max(dx, dy)
	}
	return seconds
}
