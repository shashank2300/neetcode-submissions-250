func isMatch(s string, p string) bool {
    sLen, pLen := len(s), len(p)

	memo := make(map[[2]int]bool)

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		idx := [2]int{i, j}
		if val, ok := memo[idx]; ok {
			return val
		}
		if i >= sLen && j >= pLen {
			return true
		}
		if j >= pLen {
			return false
		}

		match := i < sLen && (s[i] == p[j] || p[j] == '.')

		if j+1 < pLen && p[j+1] == '*' {
			memo[idx] = dfs(i, j+2) || (match && dfs(i+1, j))
			return memo[idx]
		}

		if match {
			memo[idx] = dfs(i+1, j+1)
			return memo[idx]
		}

		memo[idx] = false
		return false
	}

	return dfs(0, 0)
}
