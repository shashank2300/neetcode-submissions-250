func shipWithinDays(weights []int, days int) int {
	l, r := 0, 0
	for _, weight := range weights {
		if weight > l {
			l = weight
		}
		r += weight
	}
	ship_capacity := r
	for l <=r {
		m := l + (r-l)/2
		cur_load := 0
		cur_days := 0
		for _, w := range weights {
			if cur_load + w <= m {
				cur_load += w
			} else {
				cur_days++
				cur_load = w
			}
		}
		if cur_load > 0 {
			cur_days++
		}
		if cur_days <= days {
			ship_capacity = min(ship_capacity, m)
			r = m-1
		} else {
			l = m+1
		}
	}
	return ship_capacity
}
