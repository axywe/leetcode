package main

func appendCharacters(s string, t string) int {
	ptr := 0
	for i := 0; i < len(s); i++ {
		if s[i] == t[ptr] {
			ptr++
			if ptr == len(t) {
				return 0
			}
		}
	}
	return len(t) - ptr
}
