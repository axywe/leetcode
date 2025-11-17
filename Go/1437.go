package main

func kLengthApart(nums []int, k int) bool {
	c := 0
	s := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			c++
		} else {
			if s && c < k {
				return false
			}
			s = true
			c = 0
		}
	}
	return true
}
