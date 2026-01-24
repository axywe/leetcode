package main

import "slices"

func minPairSum(nums []int) int {
	slices.Sort(nums)
	max := 0
	for i := 0; i < (len(nums)-1)/2+1; i++ {
		if nums[i]+nums[len(nums)-i-1] > max {
			max = nums[i] + nums[len(nums)-i-1]
		}
	}
	return max
}
