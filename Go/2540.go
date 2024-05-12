package main

func getCommon(nums1 []int, nums2 []int) int {
	m := make(map[int]bool, len(nums1))
	for _, i := range nums1 {
		m[i] = true
	}
	for _, i := range nums2 {
		if m[i] == true {
			return i
		}
	}
	return -1
}
