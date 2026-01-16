package main

func combineRecursive(cur []int, n, level int) [][]int {
	result := [][]int{}
	i := 1
	if len(cur) != 0 {
		i = cur[len(cur)-1] + 1
	}
	if level == 0 {
		for ; i <= n; i++ {
			cp := append([]int{}, cur...)
			result = append(result, append(cp, i))
		}
		return result
	}
	for ; i <= n; i++ {
		cp := append([]int{}, cur...)
		result = append(result, combineRecursive(append(cp, i), n, level-1)...)
	}
	return result
}

func combine(n int, k int) [][]int {
	return combineRecursive([]int{}, n, k-1)
}
