func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
    var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i<0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}
	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}

	return ans
}
