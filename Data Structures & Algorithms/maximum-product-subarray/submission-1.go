func maxProduct(nums []int) int {
	n := len(nums)
    pos, neg := make([]int, n), make([]int, n)

	ans := -9999999999
	for i := range nums {
		if i == 0 {
			if nums[i] > 0 {
				pos[i] = nums[i]
				ans = pos[i]
			} else {
				neg[i] = nums[i]
				ans = neg[i]
			}

		} else {
			pos[i] = max(pos[i-1]*nums[i], neg[i-1]*nums[i], nums[i])
			neg[i] = min(neg[i-1]*nums[i], pos[i-1]*nums[i], nums[i])

			ans = max(ans, pos[i])
		}

	}
	return ans
}
