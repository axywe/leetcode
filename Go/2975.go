package main

import "sort"

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	if n == m {
		return ((m - 1) * (n - 1)) % 1000000007
	}

	hFences = append(hFences, []int{1, m}...)
	vFences = append(vFences, []int{1, n}...)
	sort.Ints(hFences)
	sort.Ints(vFences)

	fencesH := make(map[int]bool)
	for i := 1; i < len(hFences); i++ {
		for j := 0; j < i; j++ {
			fencesH[hFences[i]-hFences[j]] = true
		}
	}
	max := -1
	for i := 1; i < len(vFences); i++ {
		for j := 0; j < i; j++ {
			if fencesH[vFences[i]-vFences[j]] && max < vFences[i]-vFences[j] {
				max = vFences[i] - vFences[j]
			}
		}
	}
	if max == -1 {
		return -1
	}
	return max * max % 1000000007
}
