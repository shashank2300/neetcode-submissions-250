func minPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	dp := make([][]int, rows)

	for i := range dp {
		dp[i] = make([]int, cols)
	}

	for j := range cols {
		dp[0][j] = grid[0][j]
		if j == 0 {
			continue
		}
		dp[0][j] += dp[0][j-1]
	}
	for i := range rows {
		dp[i][0] = grid[i][0]
		if i > 0 {
			dp[i][0] += dp[i-1][0]
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[rows-1][cols-1]
}
