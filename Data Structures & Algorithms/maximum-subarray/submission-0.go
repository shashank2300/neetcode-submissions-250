func maxSubArray(nums []int) int {
    ans := -9999999999
	cur := 0

	for _, num := range nums {
		cur += num
		ans = max(ans, cur)
		if cur < 0 {
			cur = 0
		}
	}

	return ans 
}
