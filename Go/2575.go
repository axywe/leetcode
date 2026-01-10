package main

func divisibilityArray(word string, m int) []int {
	div := int64(word[0]-'0') % int64(m)
	result := make([]int, len(word))
	if div == 0 {
		result[0] = 1
	}
	for i := 1; i < len(word); i++ {
		div = (div*10 + int64(word[i]-'0')) % int64(m)
		if div == 0 {
			result[i] = 1
		}
	}
	return result
}
