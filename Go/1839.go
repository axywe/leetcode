package main

func longestBeautifulSubstring(word string) int {
    if len(word) < 5 {
        return 0
    }
    longest := 0
    cur := 0
    isB := false
    m := map[byte]int {
        'a' : 1,
        'e' : 2,
        'i' : 3,
        'o' : 4,
        'u' : 5,
    }
    if word[0] == 'a' {
        isB = true
        cur++
    }
    for i:=1;i<len(word);i++ {
        if m[word[i]] < m[word[i-1]] || m[word[i]] > m[word[i-1]]+1 {
            cur = 0
            isB = false
        }

        if word[i] == 'a' {
            isB = true
            cur++
        } else if isB == true && m[word[i]] >= m[word[i-1]] {
            cur++
        }
        if isB == true && word[i] == 'u' && cur > longest {
            longest = cur
        }
    }
    return longest
}