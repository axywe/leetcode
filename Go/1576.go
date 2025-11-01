package main

func modifyString(s string) string {
	res := []byte{}
	if s == "?" {
		return "a"
	}
	if len(s) == 1 {
		return s
	}
	if s[0] == '?' {
		if s[1] == 'a' {
			res = append(res, 'b')
		} else {
			res = append(res, 'a')
		}
	} else {
		res = append(res, s[0])
	}
	for i := 1; i < len(s)-1; i++ {
		if s[i] != '?' {
			res = append(res, s[i])
		} else {
			if res[i-1] != 'a' && s[i+1] != 'a' {
				res = append(res, 'a')
			} else if res[i-1] != 'b' && s[i+1] != 'b' {
				res = append(res, 'b')
			} else if res[i-1] != 'c' && s[i+1] != 'c' {
				res = append(res, 'c')
			}
		}
	}
	if s[len(s)-1] == '?' {
		if res[len(s)-2] == 'a' {
			res = append(res, 'b')
		} else {
			res = append(res, 'a')
		}
	} else {
		res = append(res, s[len(s)-1])
	}
	return string(res)
}
