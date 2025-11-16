package main

func numSub(s string) int {
	cur, total := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cur++
		} else {
			cur = 0
		}
		total = (total + cur) % (1e9 + 7)
	}
	return total
}
