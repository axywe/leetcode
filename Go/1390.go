package main

func sumFourDivisors(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		div := 0
		c := 0
		for j := 1; j*j <= nums[i]; j++ {
			if nums[i]%j == 0 {
				div += j
				c++
				if j*j != nums[i] {
					div += nums[i] / j
					c++
				}
			}
			if c > 4 {
				break
			}
		}
		if c == 4 {
			result += div
		}
	}
	return result
}
