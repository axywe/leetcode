package main

func maximizeSum(nums []int, k int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	result := 0
	for i := 0; i < k; i++ {
		result += max
		max++
	}
	return result
}
