func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := 0
	prevEnd := intervals[0][1]
	for _, interval := range intervals[1:] {
		start, end := interval[0], interval[1]
		if start >= prevEnd {
			prevEnd = end
		} else {
			ans++
			prevEnd = min(prevEnd, end)
		}
	}
	return ans
}
