func rob(nums []int) int {
    n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n < 3 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, n)
	copy(dp, nums)
	dp[2] += dp[0]
	for i:=3; i<n; i++ {
		dp[i] += max(dp[i-2], dp[i-3])
	}
	return max(dp[n-1], dp[n-2])
}
