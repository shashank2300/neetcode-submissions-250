func partitionLabels(s string) []int {
    lastAppearance := make(map[byte]int, 26)
	for i, c := range s {
		lastAppearance[byte(c)] = i
	}
	ans := make([]int, 0)

	for i := 0; i < len(s);  {
		end := lastAppearance[s[i]]
		for j:= i; j <= end && j < len(s); j++ {
			end = max(end, lastAppearance[s[j]])
		}
		ans = append(ans, end-i+1)
		i = end+1
	}
	return ans
}
