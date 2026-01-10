package main

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]int, len(s2)+1)
	}

	sum := 0
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = int(s1[i-1]) + sum
		sum += int(s1[i-1])
	}

	sum = 0
	for i := 1; i <= len(s2); i++ {
		dp[0][i] = int(s2[i-1]) + sum
		sum += int(s2[i-1])
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
