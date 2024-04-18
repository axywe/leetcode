package main

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, byte(s[i]))
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != byte(s[i]-1) && stack[len(stack)-1] != byte(s[i]-2) {
				return false
			} else {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
