package main

import "math"

func toBinary(num int) []int {
	res := []int{}
	for i := num; i > 0; i /= 2 {
		if i%2 == 0 {
			res = append(res, 0)
		} else {
			res = append(res, 1)
		}
	}
	return res
}

func removeLastBit(arr []int) int {
	changed := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[i-1] = 0
			changed = true
			break
		}
	}
	if changed == false {
		arr[len(arr)-1] = 0
	}
	result := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			result += int(math.Pow(2, float64(i)))
		}
	}
	return result
}

func minBitwiseArray(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 2 {
			result[i] = -1
			continue
		}
		binReverse := toBinary(nums[i])
		result[i] = removeLastBit(binReverse)
	}
	return result
}
