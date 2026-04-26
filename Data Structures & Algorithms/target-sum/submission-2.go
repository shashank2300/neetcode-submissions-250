func findTargetSumWays(nums []int, target int) int {
	total := 0 
	for _, num := range nums {
		total += num
	}
	if total < abs(target) || (total+target)%2 == 1 {
		return 0
	}
	amount := (total+target)/2
	return change(amount, nums)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
				dp[i][j] += dp[i-1][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}
