package main

func maxDotProduct(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
	}

	dp[0][0] = nums1[0] * nums2[0]

	for i := 1; i < l1; i++ {
		if nums1[i]*nums2[0] > dp[i-1][0] {
			dp[i][0] = nums1[i] * nums2[0]
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < l2; i++ {
		if nums1[0]*nums2[i] > dp[0][i-1] {
			dp[0][i] = nums1[0] * nums2[i]
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			v1 := nums1[i] * nums2[j]
			if dp[i-1][j-1] > 0 {
				v1 += dp[i-1][j-1]
			}
			v2 := dp[i-1][j]
			v3 := dp[i][j-1]
			dp[i][j] = max(v1, v2, v3)
		}
	}
	return dp[l1-1][l2-1]
}
