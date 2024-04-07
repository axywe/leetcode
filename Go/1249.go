package main

func minRemoveToMakeValid(s string) string {
	var bracket int = 0
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == '(' {
			bracket++
		} else if b[i] == ')' && bracket > 0 {
			bracket--
		} else if b[i] == ')' {
			b = append(b[:i], b[i+1:]...)
			if i-1 < -1 {
				i = -1
			} else {
				i--
			}
		}
	}
	bracket = 0
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == ')' {
			bracket++
		} else if b[i] == '(' && bracket > 0 {
			bracket--
		} else if b[i] == '(' {
			b = append(b[:i], b[i+1:]...)
		}
	}
	return string(b)
}
