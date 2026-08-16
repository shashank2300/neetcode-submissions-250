func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	isValid := func(i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		return true
	}
    var dfs func(i, j int)
	dfs = func(i, j int) {
		if !isValid(i, j) {
			return
		}
		if grid[i][j] == '0' {
			return 
		}
		grid[i][j] = '0'
		dfs(i, j+1)
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j-1)
	}
	var ans int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	fmt.Println(grid)
	return ans
}
