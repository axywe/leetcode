package main

func maxFreqSum(s string) int {
    v := make(map[byte]int, 5)
    nv := make(map[byte]int, 21)
    vowels := []byte{'a', 'e', 'i', 'o', 'u'}
    for i:=0;i<len(s);i++ {
        if slices.Contains(vowels, s[i]) {
            v[s[i]]++
        } else {
            nv[s[i]]++
        }
    }
    sum := 0
    c := 0
    for _, count := range v {
        if c < count {
            c = count
        }
    }
    sum += c
    c = 0
    for _, count := range nv {
        if c < count {
            c = count
        }
    }
    sum += c
    return sum
}