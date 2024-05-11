package main

func removeDuplicates(nums []int) int {
	current := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[current] = nums[i]
			current++
		}
	}
	return current
}
