package main

func maxDepth(s string) int {
	max := 0
	current := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			current++
		} else if s[i] == ')' {
			current--
		}
		if current > max {
			max = current
		}
	}
	return max
}
