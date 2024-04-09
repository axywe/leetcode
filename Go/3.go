package main

func lengthOfLongestSubstring(s string) int {
	size := 0
	ptr := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if val, ok := m[s[i]]; ok && val >= ptr {
			ptr = val + 1
		}
		size = max(size, i-ptr+1)
		m[s[i]] = i
	}
	return size
}
