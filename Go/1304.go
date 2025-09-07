package main

func sumZero(n int) []int {
	arr := make([]int, 0)
	if n%2 == 1 {
		arr = append(arr, 0)
		n -= 1
	}
	for i := 0; i < n/2; i++ {
		arr = append(arr, (i + 1))
		arr = append(arr, -(i + 1))
	}
	return arr
}