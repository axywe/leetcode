package main

func search(nums []int, target int) int {
	mid := len(nums) - 1
	result := 0
	switch {
	case len(nums) == 0:
		result = -1
	case nums[mid] > target:
		result = search(nums[:mid], target)
	case nums[mid] < target:
		result = search(nums[mid+1:], target)
		if result >= 0 {
			result += mid + 1
		}
	default:
		result = mid
	}

	if result == -1 {
		mid = 0
		result = 0
		switch {
		case len(nums) == 0:
			result = -1
		case nums[mid] > target:
			result = search(nums[:mid], target)
		case nums[mid] < target:
			result = search(nums[mid+1:], target)
			if result >= 0 {
				result += mid + 1
			}
		default:
			result = mid
		}
	}
	return result
}
