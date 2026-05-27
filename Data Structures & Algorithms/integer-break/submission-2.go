func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[1] = 1

	for num := 2; num <= n; num++ {
		if num == n {
			dp[num] = 0
		} else {
			dp[num] = num
		}

		for j := 1; j < num; j++ {
			if dp[j]*dp[num-j] > dp[num] {
				dp[num] = dp[j] * dp[num-j]
			}
		}
	}

	return dp[n]
}
