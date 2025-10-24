package main

func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    count := 0
    m := make(map[byte]int)
    for l, r:=0, 0; r<len(s);r++ {
        k, e := m[s[r]]
        if e && k+1 > l{
            l = k+1
        }
        m[s[r]] = r
        if r-l+1 > count {
            count = r-l+1
        }
    }
    return count
}
