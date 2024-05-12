package main

func findFinalValue(nums []int, original int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}
	for m[original] == true {
		original *= 2
	}
	return original
}
