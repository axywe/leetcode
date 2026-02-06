package main

import (
	"math"
	"sort"
)

func minRemoval(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}
	sort.Ints(nums)
	min := math.MaxInt
	for i, j := 0, 1; j < len(nums); j++ {
		if nums[i]*k < nums[j] {
			i++
		}
		if min > len(nums)-(j-i+1) {
			min = len(nums) - (j - i + 1)
		}
	}
	return min
}
