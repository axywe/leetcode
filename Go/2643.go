package main

func rowAndMaximumOnes(mat [][]int) []int {
	result := []int{0, 0}
	for i := range mat {
		count := 0
		for _, one := range mat[i] {
			if one == 1 {
				count++
			}
		}
		if count > result[1] {
			result[0] = i
			result[1] = count
		}
	}
	return result
}
