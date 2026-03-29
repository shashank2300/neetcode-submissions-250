func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	res := nums[0]
	for l <= r {
		if nums[l] < nums[r] {
			if nums[l] < res {
				res = nums[l]
				break
			}
		}
		m := l + (r-l)/2
		res = min(nums[m], res)
		if nums[l] <= nums[m] {
			l = m+1
		} else  {
			r = m-1
		}
	}
	return res
}
