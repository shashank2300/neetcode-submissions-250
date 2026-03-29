func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	for i := 0; i < len(intervals);  {
		start := intervals[i][0]
		end := intervals[i][1]
		j := i
		if i < len(intervals)-1 && intervals[i][1] >= intervals[i+1][0] {
			j++
			for ; j < len(intervals) && end >= intervals[j][0]; j++ {
				end = max(end, intervals[j][1])
			}
			j--
		}
		ans = append(ans, []int{start, end})
		i = j+1
	}

	return ans
}
