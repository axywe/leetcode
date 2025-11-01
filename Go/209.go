package main

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 1 {
		if target < nums[0] {
			return 1
		} else {
			return 0
		}
	}
	min := math.MaxInt32
	c1, c2 := 0, 1
	sum := nums[c1] + nums[c2]
	for c2 < len(nums)-1 || c1 < c2 {
		if nums[c1] >= target || nums[c2] >= target {
			return 1
		}
		if sum >= target && c2-c1+1 <= min {
			min = c2 - c1 + 1
		}
		if (c1+1 == c2 || sum < target) && c2 != len(nums)-1 {
			c2++
			sum += nums[c2]
			continue
		}
		if sum >= target || c2 == len(nums)-1 {
			sum -= nums[c1]
			c1++
		}
	}
	if min == math.MaxInt32 {
		min = 0
	}
	return min
}
