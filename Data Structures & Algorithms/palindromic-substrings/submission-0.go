func countSubstrings(s string) int {
    res := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r <len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r <len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}
