func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)

	for i := range numbers {
		if idx, ok := m[target-numbers[i]]; ok {
			return []int{idx+1, i+1}
		}
		m[numbers[i]] = i
	}

	return nil
}
