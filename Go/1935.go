package main

func canBeTypedWords(text string, brokenLetters string) int {
	flag := false
	count := 0
	for _, i := range text {
		if i == ' ' && flag == false {
			count++
		} else if i == ' ' && flag == true {
			flag = false
		}
		for _, j := range brokenLetters {
			if i == j {
				flag = true
			}
		}
	}
	if flag == false {
		count++
	}
	return count
}
