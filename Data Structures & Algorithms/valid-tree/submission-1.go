func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
        return false
    }

    adj := make(map[int][]int, 0)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make(map[int]bool)

	var dfs func(cur, parent int) bool
	dfs = func(cur, parent int) bool {
		if visited[cur] {
			return false
		}

		visited[cur] = true
		for _, neighbor := range adj[cur] {
			if neighbor == parent {
				continue
			}
			if !dfs(neighbor, cur) {
				return false
			}
		}
		return true
	}

	if !dfs(0, -1) {
		return false
	}
	return len(visited) == n
}
