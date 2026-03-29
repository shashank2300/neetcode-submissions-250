func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	rotated := append(nums[n-k:], nums[:n-k]...)
	copy(nums, rotated)
}
