package main

func firstUniqChar(s string) int {
	m := make(map[int]int, 26)
	for i, ch := range s {
		if m[int(ch)] == 0 {
			m[int(ch)] = i + 1
		} else {
			m[int(ch)] = -1
		}
	}
	min := len(s)
	for _, i := range m {
		if i != -1 && i-1 < min {
			min = i - 1
		}
	}
	if min == len(s) {
		return -1
	}
	return min
}
