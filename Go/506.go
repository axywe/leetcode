package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	resultInt := []int{}
	sorted := make([]int, len(score))
	_ = copy(sorted, score)
	sort.Sort(sort.IntSlice(sorted))

	for _, sc := range score {
		resultInt = append(resultInt, len(score)-sort.SearchInts(sorted, sc))
	}

	result := []string{}
	for _, i := range resultInt {
		switch i {
		case 1:
			result = append(result, "Gold Medal")
		case 2:
			result = append(result, "Silver Medal")
		case 3:
			result = append(result, "Bronze Medal")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
