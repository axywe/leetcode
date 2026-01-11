package main

func maximalRectangle(matrix [][]byte) int {
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			limit := len(matrix[i])
			for m := i; m < len(matrix) && matrix[m][j] != '0'; m++ {
				k := j
				for ; k < limit; k++ {
					if matrix[m][k] == '0' {
						limit = k
						break
					}
				}
				k--
				if (m-i+1)*(k-j+1) > max {
					max = (m - i + 1) * (k - j + 1)
				}
			}
		}
		if max >= len(matrix[i])*len(matrix)-i {
			return max
		}
	}
	return max
}
