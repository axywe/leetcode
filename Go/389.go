package main

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int, 26)
	for i := range s {
		m[s[i]]++
		m[t[i]]--
	}
	m[t[len(s)]]--
	for i, c := range m {
		if c != 0 {
			return i
		}
	}
	return '-'
}
