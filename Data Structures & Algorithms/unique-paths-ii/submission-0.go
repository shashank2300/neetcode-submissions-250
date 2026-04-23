func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp[m-1][n-1] = 1

	for i := m-1; i >=0; i-- {
		for j := n-1; j>=0; j-- {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i + 1 < m {
				dp[i][j] += dp[i+1][j]
			}
			if j + 1 < n {
				dp[i][j] += dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}
