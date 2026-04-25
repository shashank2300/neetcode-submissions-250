func canPartition(nums []int) bool {
    n := len(nums)
	total := 0
	for _, num := range nums {
		total += num
	}

	target := total/2

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i-1] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i-1]] + nums[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return total == dp[n][target]*2
}
