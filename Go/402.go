package main

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	for i := 0; i < len(num)-1 && k > 0; i++ {
		if num[i] > num[i+1] {
			num = num[:i] + num[i+1:]
			k--
			i -= 2
			if i == -2 {
				i = -1
			}
		}
	}
	if k > 0 {
		num = num[:len(num)-k]
	}
	for i := 0; i < len(num); i++ {
		if len(num) != 1 && num[i] == '0' {
			num = num[1:]
			i--
		} else {
			break
		}
	}
	return num
}
