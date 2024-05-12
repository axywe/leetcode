package main

func firstPalindrome(words []string) string {
	for _, word := range words {
		if len(word) == 1 {
			return word
		}
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			if word[i] != word[j] {
				break
			}
			if j-i <= 2 {
				return word
			}
		}
	}
	return ""
}
