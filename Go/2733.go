package main

func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	f, s, t := nums[0], nums[1], nums[2]
	if (f > s && f < t) || (f < s && f > t) {
		return f
	}
	if (s > t && s < f) || (s < t && s > f) {
		return s
	}
	if (t > f && t < s) || (t < f && t > s) {
		return t
	}
	return -1
}
