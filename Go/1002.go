package main

func commonChars(words []string) []string {
	m1 := make(map[rune]int)
	for _, c := range words[0] {
		m1[c]++
	}
	for _, word := range words {
		m2 := make(map[rune]int)
		for _, sym := range word {
			m2[sym]++
		}
		for k := range m1 {
			if m1[k] > m2[k] {
				m1[k] = m2[k]
			}
		}
	}
	var result []string
	for k, v := range m1 {
		for i := 0; i < v; i++ {
			result = append(result, string(k))
		}
	}
	return result
}
