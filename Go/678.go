package main

func checkValidString(s string) bool {
	counter, stars := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			counter++
		case ')':
			counter--
			if counter+stars < 0 {
				return false
			}
		case '*':
			stars++
		}
	}
	if counter-stars > 0 {
		return false
	}
	counter, stars = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			counter++
			if counter-stars > 0 {
				return false
			}
		case ')':
			counter--
		case '*':
			stars++
		}
	}
	return true
}
