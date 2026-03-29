

func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}

	tMap := make(map[byte]int)
	for i := range t {
		tMap[t[i]]++
	}

	window := make(map[byte]int)
	l, r := 0, 0
	have, need := 0, len(tMap)
	ansL, ansR := -1, -1
	ansLen := math.MaxInt32
	for ; r < len(s); r++ {
		window[s[r]]++

		if tMap[s[r]] > 0 && window[s[r]] == tMap[s[r]] {
			have++
		}
		for have == need {
			if r-l+1 < ansLen {
				ansL, ansR = l, r
				ansLen = r-l+1
			}

			window[s[l]]--
			if tMap[s[l]] > 0 && window[s[l]] < tMap[s[l]] {
				have--
			}
			l++
		}
	}

	if ansL == -1 {
		return ""
	}

	return s[ansL:ansR+1]

}
