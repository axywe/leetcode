package main

func nextBeautifulNumber(n int) int {
    arr := make([]int, 10)
    for i:=n+1;i<1000000;i++ {
        for j:=i;j>0;j/=10 {
            arr[j%10]++
        }
        isResult := true
        for j:=0;j<len(arr);j++ {
            if arr[j] != 0 && arr[j] != j {
                isResult = false
            }
            arr[j] = 0
        }
        if isResult {
            return i
        }
    }
    return 1224444
}