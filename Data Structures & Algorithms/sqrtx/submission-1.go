func mySqrt(x int) int {
	l, r := 0, (x+1)/2

	for l <= r {
		m := l + (r-l)/2

		if m*m == x {
			return m
		} else if m*m < x {
			l = m+1
		} else {
			r = m-1
		}
	}
	return l-1
}
