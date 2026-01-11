package main

func generator(s string, opened, total int) []string {
	if opened == 0 && total == 0 {
		return []string{s}
	}
	res := []string{}
	if opened != 0 {
		res = append(res, generator(s+")", opened-1, total)...)
	}
	if total != 0 {
		res = append(res, generator(s+"(", opened+1, total-1)...)
	}
	return res
}

func generateParenthesis(n int) []string {
	return generator("(", 1, n-1)
}
