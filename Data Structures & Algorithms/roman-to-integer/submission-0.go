func romanToInt(s string) int {
	d := make(map[byte]int)

	d['I'] = 1
	d['V'] = 5
	d['X'] = 10
	d['L'] = 50
	d['C'] = 100
	d['D'] = 500
	d['M'] = 1000

	tot, cur := 0, 0

	for i := range s {
		if d[s[i]] > cur {
			tot -= cur
			cur = d[s[i]]
		} else if d[s[i]] < cur {
			tot += cur
			cur = d[s[i]]
		} else {
			cur += d[s[i]]
		}
	}

	tot += cur

	return tot
}
