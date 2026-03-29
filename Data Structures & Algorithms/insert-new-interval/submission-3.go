func insert(intervals [][]int, newInterval []int) [][]int {
	// if len(intervals) == 0 {
	// 	return [][]int{newInterval}
	// }
    ans := make([][]int, 0)
	i := 0
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; {
		ans = append(ans, intervals[i])
		i++
	}
	// no overlap
	if i >= len(intervals) || newInterval[1] < intervals[i][0] {
		ans = append(ans, newInterval)
	} else {
		start, end := min(newInterval[0], intervals[i][0]), max(newInterval[1], intervals[i][1])
		i++
		for ; i < len(intervals) && end >= intervals[i][0]; i++ {
			end = max(end, intervals[i][1])
		}
		ans = append(ans, []int{start, end})
	}

	for ; i < len(intervals); i++ {
		ans = append(ans, intervals[i])
	}

	return ans
}
