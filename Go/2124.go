package main

func checkString(s string) bool {
	b := false
	for _, c := range s {
		if c == 'a' && b == true {
			return false
		}
		if c == 'b' {
			b = true
		}
	}
	return true
}
