func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)

	l, r := 1, piles[len(piles)-1]
	k := 9999999999
	for l <= r {
		m := l + (r-l)/2
		cur := 0
		for _, n := range piles {
			cur += n/m
			if n%m > 0 {
				cur++
			}
		}
		if cur <= h {
			k = min(k, m)
			r = m-1
		} else {
			l = m+1
		}
	}
	return k
}
