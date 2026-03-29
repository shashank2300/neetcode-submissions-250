func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	t0, t1, t2 := 0, 1, 1
	var t int
	for i := 3; i<=n; i++ {
		t = t2+t1+t0
		t2, t1, t0 = t, t2, t1
	}

	return t2
}
