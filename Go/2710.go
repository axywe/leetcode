package main

func removeTrailingZeros(num string) string {
    count := len(num)
    for i := len(num)-1;i>=0;i-- {
        if num[i] == '0' {
            count--
        } else {
            break
        }
    }
    result := make([]byte, count)
    for i := 0; i < count; i++ {
        result[i] = num[i]
    }
    return string(result)
}