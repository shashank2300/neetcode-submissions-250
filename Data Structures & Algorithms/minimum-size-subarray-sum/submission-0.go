func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	ans, cur := 999999999, 0
	for r < len(nums) {
		cur += nums[r]

		for cur >= target {
			ans = min(ans, r-l+1)
			cur -= nums[l]
			l++
		}
		r++
	}
	if ans == 999999999 {
		return 0
	}
	return ans
}
