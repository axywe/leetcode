package main

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}
	result := string(s[0])
	for i := 1; i < len(s)-1; i++ {
		if s[i-1] == s[i+1] {
			for j, k := i-1, i+1; j >= 0 && k < len(s); {
				if s[j] == s[k] {
					if len(result) < k-j+1 {
						result = s[j : k+1]
					}
					j--
					k++
				} else {
					break
				}
			}
		}

		if s[i-1] == s[i] {
			for j, k := i-1, i; j >= 0 && k < len(s); {
				if s[j] == s[k] {
					if len(result) < k-j+1 {
						result = s[j : k+1]
					}
					j--
					k++
				} else {
					break
				}
			}
		}
	}
	if s[len(s)-2] == s[len(s)-1] && len(result) == 1 {
		result = s[len(s)-2 : len(s)]
	}
	return result
}
