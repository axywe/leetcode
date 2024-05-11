package main

func titleToNumber(columnTitle string) int {
	result := 0
	for _, ch := range columnTitle {
		result = result*26 + int(byte(ch)-64)
	}
	return result
}
