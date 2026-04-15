func canJump(nums []int) bool {
    n := len(nums)
	// dp := make([]bool, n)

	// dp[0] = true
	next := nums[0]
	for i := 1; i < n; i++ {
		if i > next {
			return false
		}
		next = max(next, i + nums[i])
	}
	return true
}
