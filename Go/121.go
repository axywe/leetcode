package main

func maxProfit(prices []int) int {
    min, max := prices[0], 0
    for i:=0;i<len(prices);i++ {
        if prices[i] < min {
            min = prices[i]
        }
        if prices[i]-min > max {
            max = prices[i]-min
        }
    }
    return max
}
