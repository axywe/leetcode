package main

func moveZeroes(nums []int) {
	p := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[p] = nums[i]
			p++
		}
	}
	for ; p < len(nums); p++ {
		nums[p] = 0
	}
}
