func rob(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	return max(rob_helper(nums[1:]), rob_helper(nums[:len(nums)-1]))
}

func rob_helper(nums []int) int {
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