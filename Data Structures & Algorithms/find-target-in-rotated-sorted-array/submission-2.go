func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			return m
		} else if nums[l] < nums[r] {
			if nums[m] > target {
				r = m-1
			} else {
				l = m+1
			}
		} else if nums[l] <= nums[m] {
			if target < nums[m] && target >= nums[l] {
				r = m-1
			} else {
				l = m+1
			}
		} else {
			if target > nums[m] && target <= nums[r] {
				l = m+1
			} else {
				r = m-1
			}
		}
	}
	return -1
}
