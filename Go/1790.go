package main

func areAlmostEqual(s1 string, s2 string) bool {
    var l byte
    count := 0
    sum1, sum2 := 0, 0
    for i:=0;i<len(s1);i++ {
        sum1 += int(s1[i])
        sum2 += int(s2[i])
        if count == 0 && s1[i] != s2[i] {
            l = s1[i]
            count++
        } else if count == 1 && s1[i] != s2[i] {
            if s2[i] != l {
                return false
            }
            count++
        } else if s1[i] != s2[i] {
            return false
        }
    }
    if sum1 != sum2 {
        return false
    }
    return true
}