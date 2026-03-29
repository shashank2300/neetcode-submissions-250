func findMin(nums []int) int {
	// l, r := 0, len(nums)-1

	// res := nums[0]
	// for l <= r {
	// 	if nums[l] < nums[r] {
	// 		if nums[l] < res {
	// 			res = nums[l]
	// 			break
	// 		}
	// 	}
	// 	m := l + (r-l)/2
	// 	res = min(nums[m], res)
	// 	if nums[l] <= nums[m] {
	// 		l = m+1
	// 	} else  {
	// 		r = m-1
	// 	}
	// }
	// return res
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l) / 2

		if nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}
