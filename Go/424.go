package main

func characterReplacement(s string, k int) int {
    count := 0
    m := make(map[byte]int)
    for l, r := 0, 0;r<len(s);r++ {
        sum := 0
        largest := 0
        m[s[r]]++
        for _, val := range m {
            if val > largest {
                largest = val
            }
            sum += val
        }

        if sum - largest <= k && sum > count {
            count = sum
        }
        if sum - largest > k {
            m[s[l]]--
            l++
        }
    }
    return count
}
