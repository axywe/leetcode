package main

func myAtoi(s string) int {
	out := 0
	negative := false
	positive := false
	start := false
	for _, c := range s {
		if c == 45 && start == false && negative == false && positive == false {
			negative = true
			start = true
		} else if c == 43 && start == false && negative == false && positive == false {
			positive = true
			start = true
		} else if c >= 48 && c <= 58 {
			if out > 214748364 {
				out = 2147483647
				if negative == true {
					out = -2147483648
					positive = true
				}
				break
			} else if out == 214748364 {
				if c == 56 || c == 57 || c == 48 {
					out = 2147483647
					if negative == true {
						out = -2147483648
						positive = true
					}
					break
				}
			}
			out = out*10 + (int(c) - 48)
			start = true
		} else if start == true || c != 32 {
			break
		}
	}
	if negative == true && positive != true {
		out *= -1
	}
	return out
}
