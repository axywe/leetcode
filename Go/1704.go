package main

func halvesAreAlike(s string) bool {
	vowelA, vowelB := 0, 0
	for p1, p2 := 0, len(s)-1; p1 < p2; p1, p2 = p1+1, p2-1 {
		switch s[p1] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelA++
		}
		switch s[p2] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelB++
		}
	}
	return vowelA == vowelB
}
