func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	l1, l2 := len(str1), len(str2)
	n := l2
	for n >= 1 {
		if l1 % n == 0 {
			if divisor(str2, str2[:n]) && divisor(str1, str2[:n]) {
				return str2[:n]
			}
		}
		factor := l2/n
		factor++
		for factor <= l2 {
			if l2 % factor == 0 {
				break
			}
			factor++
		}
		n = l2/factor
	}
	return ""
}

func divisor(s, divisor string) bool {
	if len(s) % len(divisor) != 0 {
		return false
	}

	m, n := len(s), len(divisor)

	for i := 0; i < m; i+= n {
		if s[i:i+n] != divisor {
			return false
		}
	}
	return true
}
