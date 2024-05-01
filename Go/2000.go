package main

func reversePrefix(word string, ch byte) string {
	result := ""
	for i, c := range word {
		if byte(c) == ch {
			for j := 0; j < len(word); j++ {
				if j <= i {
					result += string(word[i-j])
				} else {
					result += string(word[j])
				}
			}
			return result
		}
	}
	return word
}
