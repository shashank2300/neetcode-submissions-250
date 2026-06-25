func groupAnagrams(strs []string) [][]string {
    n := len(strs)

    counts := make(map[string][]string, n)

    for _, s := range strs {
        count := [26]int{}
        for _, c := range s {
            count[c-'a']++
        }
        temp := frequencyToString(count)
        counts[temp] = append(counts[temp], s)
    }
    ans := [][]string{}
    for _, v := range counts {
        ans = append(ans, v)
    }
    return ans
}

func frequencyToString(freq [26]int) string {
	var builder strings.Builder
	for i, count := range freq {
		if count > 0 {
			builder.WriteString(fmt.Sprintf("%d%c", count, 'a'+rune(i)))
		}
	}
	return builder.String()
}