package main

func findClosest(x int, y int, z int) int {
    if math.Abs(float64(x - z)) > math.Abs(float64(y - z)) {
        return 2
    } else if math.Abs(float64(x - z)) < math.Abs(float64(y - z)) {
        return 1
    } else {
        return 0
    }
}