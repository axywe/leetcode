package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for _, ch := range magazine {
		m[byte(ch)]++
	}
	for _, ch := range ransomNote {
		m[byte(ch)]--
		if m[byte(ch)] < 0 {
			return false
		}
	}
	return true
}
