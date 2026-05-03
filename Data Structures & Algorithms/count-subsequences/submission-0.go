func numDistinct(s string, t string) int {
	rows, cols := len(s), len(t)
    dp := make([][]int, rows+1)
    for i := range dp {
        dp[i] = make([]int, cols+1)
    }

    for i := 0; i <= rows; i++ {
        dp[i][cols] = 1
    }

    for i := rows - 1; i >= 0; i-- {
        for j := cols - 1; j >= 0; j-- {
            dp[i][j] = dp[i+1][j]
            if s[i] == t[j] {
                dp[i][j] += dp[i+1][j+1]
            }
        }
    }

    return dp[0][0]
}