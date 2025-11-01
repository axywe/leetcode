package main

func canJump(nums []int) bool {
	c := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < c {
			c++
		} else {
			c = 1
		}
	}
	return c == 1
}
