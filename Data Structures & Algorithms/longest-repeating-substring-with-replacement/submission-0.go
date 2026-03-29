func characterReplacement(s string, k int) int {
	m := make(map[byte]int)

	res, l, maxFreq := 0, 0, 0

	for r :=0; r < len(s); r++ {
		m[s[r]]++
		if m[s[r]] > maxFreq {
			maxFreq = m[s[r]]
		}

		for (r-l+1) - maxFreq > k {
			m[s[l]]--
			l++
		}

		if r-l+1 > res {
			res = r-l+1
		}
		
	}
	return res
}
