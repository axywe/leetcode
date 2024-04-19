package main

func isMonotonic(nums []int) bool {
	status := 0 // 0 - waiting, 1 - increase, 2 - decrease
	for i := 0; i < len(nums)-1; i++ {
		if (status == 0 || status == 1) && nums[i] <= nums[i+1] {
			if status == 0 && nums[i] < nums[i+1] {
				status = 1
			}
		} else if (status == 0 || status == 2) && nums[i] >= nums[i+1] {
			if status == 0 && nums[i] > nums[i+1] {
				status = 2
			}
		} else {
			return false
		}
	}
	return true
}
