package main

func removeDuplicates(nums []int) int {
	cur := 0
	flag := false
	for i := 1; i < len(nums); i++ {
		if !flag && nums[cur] == nums[i] {
			flag = true
			cur++
			nums[cur] = nums[i]
		}
		if nums[cur] != nums[i] {
			cur++
			nums[cur] = nums[i]
			flag = false
		}
	}
	return cur + 1
}
