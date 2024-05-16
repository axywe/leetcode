package main

func commonFactors(a int, b int) int {
	factors := 0
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			factors++
		}
	}
	return factors
}
