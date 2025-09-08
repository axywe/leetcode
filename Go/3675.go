package main

func minOperations(s string) int {
    m := make(map[byte]int, 26)
    for i := 0; i < len(s); i++ {
        m[s[i]]++
    }
    for i := byte('b'); i <= byte('z'); i++ {
        if m[i] != 0 {
            return int('z' - i + 1)
        }
    }
    return 0
}©leetcode