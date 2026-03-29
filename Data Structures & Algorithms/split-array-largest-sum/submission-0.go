func splitArray(nums []int, k int) int {
	canSplit := func(maxSum int) bool {
		arrCount := 1
		cur := 0
		for _, num := range nums {
			cur += num
			if cur > maxSum {
				arrCount++
				if arrCount > k {
					return false
				}
				cur = num
			}
		}
		return true
	}

	l, r := 0, 0
	for _, num := range nums {
		if num > l {
			l = num
		}
		r += num
	}

	res := r
	for l <= r {
		m := l + (r-l)/2
		if canSplit(m) {
			res = m
			r = m-1
		} else {
			l = m + 1
		}
	}
	return res
}
