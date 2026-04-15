func lengthOfLIS(nums []int) int {
    n := len(nums)
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}

	ans := 0
	for i := n-1; i >=0 ; i-- {
		for j := i+1; j < n; j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		ans = max(ans, dp[i])
	}

	return ans
}
