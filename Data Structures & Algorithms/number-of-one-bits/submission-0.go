func hammingWeight(n int) int {
	binaryStr := strconv.FormatInt(int64(n), 2)
	
	ans := 0
	for _, c := range binaryStr {
		if c == '1' {
			ans++
		}
	}
	return ans
}
