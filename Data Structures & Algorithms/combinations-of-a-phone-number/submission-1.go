func letterCombinations(digits string) []string {
    n := len(digits)
    if n == 0 {
        return []string{}
    }
	dict := make(map[int][]byte)

    dict[2] = []byte{'a', 'b', 'c'}
    dict[3] = []byte{'d', 'e', 'f'}
    dict[4] = []byte{'g', 'h', 'i'}
    dict[5] = []byte{'j', 'k', 'l'}
    dict[6] = []byte{'m', 'n', 'o'}
    dict[7] = []byte{'p', 'q', 'r', 's'}
    dict[8] = []byte{'t', 'u', 'v'}
    dict[9] = []byte{'w','x', 'y', 'z'}


    ans := make([]string, 0)
    var recurse func(i int, cur []byte)
    recurse = func(i int, cur []byte) {
        if i >= n {
            ans = append(ans, string(append([]byte{}, cur...)))
            return
        }
        idx := int(digits[i]-'0')
        for j := range dict[idx] {
            cur = append(cur, dict[idx][j])
            recurse(i+1, cur)
            cur = cur[:len(cur)-1]
        }
    }
    recurse(0, []byte{})
    return ans
}
