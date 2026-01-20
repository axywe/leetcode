package main

func minBitwiseArray(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 1; j < nums[i]; j++ {
			if (j | (j + 1)) == nums[i] {
				result[i] = j
				break
			} else {
				result[i] = -1
			}
		}
	}
	return result
}
