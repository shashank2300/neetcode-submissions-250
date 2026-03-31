func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] < 0 {
			nums[i] = 0
		}
	}

	for i := range nums {
		val := nums[i]
		if val < 0 {
			val = -val
		}
		if val > 0 && val <= n {
			if nums[val-1] > 0 {
				nums[val-1] *= -1
			} else if nums[val-1] == 0 {
				nums[val-1] = -(n+1)
			}
		}
	}
	for i := range nums {
		if nums[i] >= 0 {
			return i+1
		}
	}
	return n+1
}
