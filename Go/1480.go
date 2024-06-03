package main

func runningSum(nums []int) []int {
	sum := 0
	for i, num := range nums {
		sum += num
		nums[i] = sum
	}
	return nums
}
