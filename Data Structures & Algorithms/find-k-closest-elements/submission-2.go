func findClosestElements(nums []int, k int, x int) []int {
	l, r := 0, k-1
	diff := 9999999999
	cur := 0
	for i := 0; i <= r-1; i++ {
		cur += abs(nums[i]-x)
	}
	ansIdx := 0
	for r < len(nums) {
		cur += abs(nums[r]-x)
		if cur < diff {
			ansIdx = l
			diff = cur
		}
		cur -= abs(nums[l]-x)
		l++
		r++
		
	}
	return nums[ansIdx:ansIdx+k]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
