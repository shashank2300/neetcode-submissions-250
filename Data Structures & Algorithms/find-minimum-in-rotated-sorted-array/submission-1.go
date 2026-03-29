func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	res := nums[0]
	for l <= r {
		m := l + (r-l)/2
		res = min(nums[m], res)
		if nums[l] < nums[m] && nums[m] < nums[r] {
			r = m-1
		} else if nums[l] < nums[m] && nums[m] > nums[r] {
			l = m+1
		} else if nums[l] > nums[m] {
			r = m-1
		} else {
			l = m+1
		}
	}
	return res
}
