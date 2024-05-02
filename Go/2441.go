package main

func findMaxK(nums []int) int {
	num := make(map[int]int, 1001)
	//1 - negative, 2 - positive, 3 - both
	max := -1
	for _, i := range nums {
		if i < 0 && num[-i] != 2 {
			num[-i] = 1
		} else if i > 0 && num[i] != 1 {
			num[i] = 2
		} else {
			if i > max {
				max = i
			} else if -i > max {
				max = -i
			}
		}
	}
	return max
}
