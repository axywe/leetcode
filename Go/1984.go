package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	result := math.MaxInt
	for i := k - 1; i < len(nums); i++ {
		if result > nums[i]-nums[i-k+1] {
			result = nums[i] - nums[i-k+1]
		}
	}
	return result
}
