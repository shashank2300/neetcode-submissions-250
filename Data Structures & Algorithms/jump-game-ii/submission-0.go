func jump(nums []int) int {
	n := len(nums)
    jumps := make([]int, n)
	for i := range jumps {
		jumps[i] = 999999999
	}
	jumps[n-1] = 0

	for i := n-2; i >= 0; i-- {
		for j := i+1; j < n && j <= i+nums[i]; j++ {
			jumps[i] = min(jumps[i], jumps[j]+1)
		}
	}

	return jumps[0]
}
