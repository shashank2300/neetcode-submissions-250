func maxArea(heights []int) int {
	l, r := 0, len(heights)-1

	ans := 0

	for l < r {
		ans = max(ans, (r-l)*min(heights[l], heights[r]))

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
