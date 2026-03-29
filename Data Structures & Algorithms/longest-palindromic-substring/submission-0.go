func longestPalindrome(s string) string {
    n := len(s)
	dp := make([][]bool, n)

	resIdx, resLen := 0, 0
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := n-1; i >=0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if resLen < (j-i+1) {
					resIdx = i
					resLen = j-i+1
				}
			}
		}
	}
	return s[resIdx:resIdx+resLen]
}
