func mergeAlternately(word1 string, word2 string) string {
	f, s := 0, 0
	var ans strings.Builder

	for f < len(word1) && s < len(word2) {
		ans.WriteByte(word1[f])
		ans.WriteByte(word2[s])
		f++
		s++
	}

	for f < len(word1) {
		ans.WriteByte(word1[f])
		f++
	}

	for s < len(word2) {
		ans.WriteByte(word2[s])
		s++
	}

	return ans.String()
}
