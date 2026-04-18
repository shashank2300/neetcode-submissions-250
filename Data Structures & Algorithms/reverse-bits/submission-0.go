func reverseBits(n int) int {
	ans := 0
	for i := range 32 {
		ans ^= (n&1)<<(31-i)
		n >>= 1
	}
	return ans
}
