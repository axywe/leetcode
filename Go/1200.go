package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	res := [][]int{}
	sort.Ints(arr)
	min := math.MaxInt
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			min = arr[i] - arr[i-1]
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
