package main

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	p := 0
	for _, c := range t {
		if c == rune(s[p]) {
			p++
			if p == len(s) {
				return true
			}
		}
	}
	return false
}
