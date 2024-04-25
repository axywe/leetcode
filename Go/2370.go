package main

func longestIdealString(s string, k int) int {
	dp := [150]int{0}
	res := 0
	for _, c := range s {
		i := int(c)
		for j := i - k; j <= i+k; j++ {
			if dp[i] < dp[j] {
				dp[i] = dp[j]
			}
		}
		dp[i]++
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
