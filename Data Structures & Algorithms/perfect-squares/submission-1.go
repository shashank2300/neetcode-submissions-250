func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := range dp {
		dp[i] = i
	}
	
	for target := 2; target <= n; target++ {
		for i := 1; i*i <= target; i++ {
			dp[target] = min(dp[target], 1+dp[target-i*i])
		}
	}

	return dp[n]
}
