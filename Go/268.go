package main

func missingNumber(nums []int) int {
	sum := 0
	real := 0
	for i, j := range nums {
		sum += i + 1
		real += j
	}
	return sum - real
}
