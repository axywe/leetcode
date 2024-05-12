package main

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool, len(nums1))
	for _, i := range nums1 {
		m[i] = true
	}
	result := []int{}
	for _, i := range nums2 {
		if m[i] == true {
			result = append(result, i)
			m[i] = false
		}
	}
	return result
}
