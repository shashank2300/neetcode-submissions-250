func integerBreak(n int) int {
	if n <= 3 {
		return n-1
	}
	dp := make([]int, n+1)

	dp[1] = 1


	for num := 2; num <= n; num++ {
		dp[num] = num

		for j := 1; j < num; j++ {
			if dp[j]*dp[num-j] > dp[num] {
				dp[num] = dp[j] * dp[num-j]
			}
		}
	}

	return dp[n]
}
