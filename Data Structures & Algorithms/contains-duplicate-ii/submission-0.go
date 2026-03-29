func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)

	l, r := 0, 0
	for r < len(nums) {
		if r-l > k {
			m[nums[l]] = false
			l++
		}
		if m[nums[r]] {
			return true
		}
		m[nums[r]] = true
		r++
	}
	return false
}
