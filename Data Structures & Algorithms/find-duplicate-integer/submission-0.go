func findDuplicate(nums []int) int {
    // n := len(nums)
	m := make(map[int]bool)

	for _, num := range nums {
		if m[num] == true {
			return num
		}
		m[num] = true
	}

	return -1
}
