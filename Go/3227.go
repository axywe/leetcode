package main

func doesAliceWin(s string) bool {
	count := 0
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for i := 0; i < len(s); i++ {
		if slices.Contains(vowels, s[i]) {
			count++
		}
	}
	if count == 0 {
		return false
	}
	return true
}