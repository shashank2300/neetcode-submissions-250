func myPow(x float64, n int) float64 {
    var ans float64 = 1
	for _ = range abs(n) {
		ans *= x
	}
	if n < 0 {
		return 1/ans
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
