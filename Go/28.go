package main

func strStr(haystack string, needle string) int {
	cur := 0
	for i := 0; i < len(haystack); i++ {
		if needle[cur] == haystack[i] {
			cur++
		} else {
			i = i - cur
			cur = 0
		}
		if cur == len(needle) {
			return i - cur + 1
		}
	}
	return -1
}
