func countComponents(n int, edges [][]int) int {
    adj := make(map[int][]int, 0)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make(map[int]bool)

	var dfs func(cur int)
	dfs = func(cur int) {
		if visited[cur] {
			return
		}

		visited[cur] = true
		for _, neighbor := range adj[cur] {
			if visited[neighbor] {
				continue
			}
			dfs(neighbor)
		}
	}

	ans := 0
	for i := range n {
		if !visited[i] {
			ans++
			dfs(i)
		}
	}
	return ans
}
