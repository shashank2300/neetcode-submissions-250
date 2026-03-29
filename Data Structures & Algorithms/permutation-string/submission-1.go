func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	d := make([]int, 26)
	cur := make([]int, 26)
	l, r := 0, 0

	for r < len(s1) {
		d[rune(s1[r])-'a']++
		cur[rune(s2[r])-'a']++
		r++
	}
	equal := func(a, b []int) bool {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	if equal(d, cur) {
		return true
	}
	for r < len(s2) {
		cur[rune(s2[l])-'a']--
		cur[rune(s2[r])-'a']++
		l++
		r++
		if equal(d, cur) {
			return true
		}
	}

	return false

}
