package main

func totalMoney(n int) int {
    weeks := n/7
    result := 0
    for i:=0;i<weeks;i++ {
        result += 28 + i*7
    }
    for i:=1;i<=n%7;i++ {
        result += weeks + i
    }
    return result
}