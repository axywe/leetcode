package main

func rearrangeCharacters(s string, target string) int {
	m := make(map[rune]int, len(s))
	count := 0
	for _, ch := range s {
		m[ch]++
		if m[ch] > count {
			count = m[ch]
		}
	}
	targetMap := make(map[rune]int, len(target))
	for _, ch := range target {
		targetMap[ch]++
		if m[ch]/targetMap[ch] < count {
			count = m[ch] / targetMap[ch]
		}
	}
	return count
}
