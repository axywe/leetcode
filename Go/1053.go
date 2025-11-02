package main

func prevPermOpt1(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	i := len(arr) - 2
	for ; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			break
		}
	}
	if i == -1 {
		return arr
	}
	num1 := i
	i++
	latest := i
	for ; i < len(arr); i++ {
		if arr[num1] > arr[i] && arr[i] > arr[latest] {
			latest = i
		}
	}
	tmp := arr[num1]
	arr[num1] = arr[latest]
	arr[latest] = tmp
	return arr
}
