func change(amount int, coins []int) int {
    dp := make([][]int, len(coins)+1)

	for i := range dp {
		dp[i] = make([]int, amount + 1)
	}
	
	for i := 0; i <= len(coins); i++ {
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
				dp[i][j] += dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}
