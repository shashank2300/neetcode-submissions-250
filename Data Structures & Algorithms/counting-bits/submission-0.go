func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := range n+1 {
		ans[i] = hammingWeight(i)
	}
	return ans
}

func hammingWeight(n int) int {
	ans := 0

	for i :=0; i < 32; i++ {
		ans += (n>>i) & 1
	}
	return ans
}

