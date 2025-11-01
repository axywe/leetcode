package main

import "math"

func isCloser(num1, num2, target int) bool {
	if math.Abs(float64(target-num1)) < math.Abs(float64(target-num2)) {
		return true
	}
	return false
}

func threeSumClosest(nums []int, target int) int {
	m := make(map[int][2]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			m[nums[i]+nums[j]] = [2]int{i, j}
		}
	}

	closest := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		for k, v := range m {
			if i != v[0] && i != v[1] && isCloser(nums[i]+k, closest, target) {
				closest = nums[i] + k
			}
		}
	}
	return closest
}
