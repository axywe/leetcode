package main

import "slices"

func maximumHappinessSum(happiness []int, k int) int64 {
	slices.Sort(happiness)
	var result int64
	for i := 0; i < k; i++ {
		if int64(happiness[len(happiness)-i-1]-i) > 0 {
			result += int64(happiness[len(happiness)-i-1] - i)
		}
	}
	return result
}
