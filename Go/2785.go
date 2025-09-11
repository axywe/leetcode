package main

func sortVowels(s string) string {
	vovels := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			vovels[s[i]]++
		}
	}
	result := make([]byte, 0)
	arr := []byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			for j := 0; j < len(arr); j++ {
				if vovels[arr[j]] != 0 {
					result = append(result, arr[j])
					vovels[arr[j]]--
					break
				}
			}
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}