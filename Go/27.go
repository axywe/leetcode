package main

func removeElement(nums []int, val int) int {
	index := 0
	for i, j := range nums {
		if j != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
