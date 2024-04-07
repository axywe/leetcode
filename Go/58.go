package main

func lengthOfLastWord(s string) int {
	count := 0
	flag := false
	for i := 0; i < len(s); i++ {
		if s[i] != 32 {
			if flag == true {
				count = 0
				flag = false
			}
			count++
		} else {
			flag = true
		}
	}
	return count
}
