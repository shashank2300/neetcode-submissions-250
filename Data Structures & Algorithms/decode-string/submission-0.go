func decodeString(s string) string {
	counts, chars := make([]int, 0), make([]string, 0)

	cur := ""
	k := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			k = k*10 + int(c-'0')
		} else if c == '[' {
			chars = append(chars, cur)
			counts = append(counts, k)
			cur = ""
			k = 0
		} else if c == ']' {
			temp := cur
			cur = chars[len(chars)-1] // the string before current string
			// this is equivalent to prev then adding count*cur
			chars = chars[:len(chars)-1]
			count := counts[len(counts)-1]
			counts = counts[:len(counts)-1]
			for _ = range count {
				cur += temp
			}
		} else {
			cur += string(c)
		}
	}
	return cur
}
