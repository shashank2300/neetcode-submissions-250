func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MinInt32
	}

	dp[n] = 0

	for i := n-1; i >= 0; i-- {
		total := 0
		for j := i; j < min(i+3, n); j++ {
			total += stoneValue[j]
			dp[i] = max(dp[i], total-dp[j+1])
		}
	}

	if dp[0] == 0 {
		return "Tie"
	} else if dp[0] > 0 {
		return "Alice"
	}

	return "Bob"
}
