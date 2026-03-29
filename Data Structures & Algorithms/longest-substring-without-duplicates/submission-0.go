func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)

	ans := 0
	l, r := 0, 0
	for r < len(s) {
		for m[s[r]] {
			m[s[l]] = false
			l++
		}
		ans = max(ans, r-l+1)
		m[s[r]] = true
		r++
	}
	return ans
}
