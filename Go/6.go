package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	matrix := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]byte, 0)
	}

	r, step := 0, -1
	for i := 0; i < len(s); i++ {
		matrix[r] = append(matrix[r], s[i])
		if r == numRows-1 || r == 0 {
			step *= -1
		}
		r += step
	}
	result := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		result = append(result, matrix[i]...)
	}
	return string(result)
}
