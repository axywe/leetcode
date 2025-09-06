package main

func letterCombinations(digits string) []string {
	m := make(map[int]string, 10)
	m[2] = "abc"
	m[3] = "def"
	m[4] = "ghi"
	m[5] = "jkl"
	m[6] = "mno"
	m[7] = "pqrs"
	m[8] = "tuv"
	m[9] = "wxyz"
	res := make([]string, 0)
    if len(digits) == 0 {
        return res
    }
    var backtrack func (combination string, nextDigits string)
    backtrack = func (combination string, nextDigits string) {
        if len(nextDigits) == 0 {
            res = append(res, combination)
            return
        }
        next := ""
        for i:=1;i<len(nextDigits);i++ {
            next += string(nextDigits[i])
        }

        for i := 0; i < len(m[int(nextDigits[0])-48]); i++ {
            backtrack(combination+string(m[int(nextDigits[0])-48][i]), next)
        }
    }
    backtrack("", digits)
	return res
}
