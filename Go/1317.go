package main

func getNoZeroIntegers(n int) []int {
    result := []int{0, 0}
    if n == 2 {
        return []int{1, 1}
    }
    for ; ; {
        rand := rand.Intn(n-1)
        str := strconv.Itoa(rand)
        flag := false
        for i := 0; i < len(str); i++ {
            if str[i] == '0' {
                flag = true
                break
            }
        }
        rand2 := n - rand
        str2 := strconv.Itoa(rand2)
        for i := 0; i < len(str2); i++ {
            if str2[i] == '0' {
                flag = true
                break
            }
        }
        if flag == false {
            result[0] = rand
            result[1] = rand2
            break
        }
    } 
    return result
}