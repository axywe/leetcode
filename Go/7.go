package main

func reverse(x int) int {
	var result int = 0
	var negative bool = false
	if x < 0 {
		x *= -1
		negative = true
	}
	for x > 0 {
		if result > 214748364 {
			return 0
		}
		result = result*10 + x%10
		x = x / 10
	}
	if negative {
		result *= -1
	}
	return result
}
