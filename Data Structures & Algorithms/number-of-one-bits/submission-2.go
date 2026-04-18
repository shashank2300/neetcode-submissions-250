func hammingWeight(n int) int {
	ans := 0

	for i :=0; i < 32; i++ {
		ans += (n>>i) & 1
	}

	return ans
}
