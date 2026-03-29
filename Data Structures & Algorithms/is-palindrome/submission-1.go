func isPalindrome(s string) bool {
	// fmt.Println("started")
	s = strings.ToLower(s)
	// fmt.Println("started")
	l, r := 0, len(s)-1
	for !isAlphaNumeric(s[l]) && l < r {
		l++
	}
	for !isAlphaNumeric(s[r]) && l < r {
		r--
	}
	
	for l < r {
		if s[r] != s[l] {
			fmt.Println(string(s[r]), string(s[l]))
			return false
		}
		l++
		r--
		for !isAlphaNumeric(s[l]) && l <= r {
			l++
		}
		for !isAlphaNumeric(s[r]) && l <= r {
			r--
		}
	}
	return true
}

func isAlphaNumeric(r byte) bool {
	return unicode.IsLetter(rune(r)) || unicode.IsDigit(rune(r))
}
