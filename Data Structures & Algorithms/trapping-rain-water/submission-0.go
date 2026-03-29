func trap(height []int) int {
	n := len(height)
	left, right := make([]int, n), make([]int, n)

	ans := 0
	for i := range height {
		if i == 0 {
			continue
		}
		left[i] = max(left[i-1], height[i-1])
		right[n-1-i] = max(right[n-i], height[n-i])
	}
	for i := range height {
		trapped_water := min(left[i], right[i])-height[i]
		if trapped_water > 0 {
			ans += trapped_water
		}
	}
	return ans
}
