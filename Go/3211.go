package main

func generator(n int, str string) []string {
	if n == 0 {
		if len(str) != 0 && str[len(str)-1] == '0' {
			return []string{str + "1"}
		} else {
			return []string{str + "0", str + "1"}
		}
	}
	res := []string{}
	if len(str) != 0 && str[len(str)-1] == '0' {
		res = append(res, generator(n-1, str+"1")...)
	} else {
		res = append(res, generator(n-1, str+"1")...)
		res = append(res, generator(n-1, str+"0")...)
	}
	return res
}

func validStrings(n int) []string {
	return generator(n-1, "")
}
