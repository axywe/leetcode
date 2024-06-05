package main

func longestPalindrome(s string) int {
	m := make(map[rune]int, 52)
	for _, ch := range s {
		m[ch]++
	}
	plus := 0
	sum := 0
	for _, i := range m {
		if i%2 == 1 {
			plus = 1
		}
		sum += i - (i % 2)
	}
	return sum + plus
}
