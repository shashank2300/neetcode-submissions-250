func longestConsecutive(nums []int) int {
	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = false
	}

	longest := 0
	for num := range m {
		if _, ok := m[num-1]; !ok {
			length := 1
			for {
				if _, ok = m[num+length]; ok {
					length++
				} else {
					break
				}
			}
			longest = max(longest, length)
		}
	}
	return longest
}
