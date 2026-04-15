func wordBreak(s string, wordDict []string) bool {
	n := len(s)
    dp := make([]bool, n+1)
	dp[n] = true

	for i := n-1; i >= 0; i-- {
		for _, w := range wordDict {
			if i+len(w) <= len(s) && s[i:i+len(w)] == w {
				dp[i] = dp[i+len(w)]
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[0]
}
