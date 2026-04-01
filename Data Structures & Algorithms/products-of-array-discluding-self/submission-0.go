func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre, post := make([]int, n), make([]int, n)

	for i := range nums {
		if i == 0 {
			pre[0] = nums[0]
			post[n-i-1] = nums[n-i-1]
			continue
		}
		pre[i] = pre[i-1]*nums[i]
		post[n-i-1] = post[n-i]*nums[n-i-1]
	}

	ans := make([]int, n)
	for i := range nums {
		if i == 0 {
			ans[i] = post[i+1]
		} else if i == n-1 {
			ans[i] = pre[i-1]
		} else {
			ans[i] = pre[i-1]*post[i+1]
		}
	}

	return ans
}
