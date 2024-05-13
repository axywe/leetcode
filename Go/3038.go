package main

func maxOperations(nums []int) int {
	operations := 1
	score := nums[0] + nums[1]
	for i := 2; i < len(nums); i += 2 {
		if i+1 < len(nums) {
			if nums[i]+nums[i+1] == score {
				operations++
			} else {
				return operations
			}
		}
	}
	return operations
}
