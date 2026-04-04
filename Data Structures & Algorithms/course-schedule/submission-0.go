func canFinish(numCourses int, prerequisites [][]int) bool {
    adj := make([][]int, numCourses)

	for _, prereq := range prerequisites {
		adj[prereq[0]] = append(adj[prereq[0]], prereq[1])
	}
	done := make([]bool, numCourses)
	for i := range numCourses {
		if len(adj[i]) == 0 {
			done[i] = true
		}
	}

	var visited = make(map[int] bool)
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if done[course] {
			return true
		}
		if visited[course] {
			return false
		}

		visited[course] = true
		for _, prereq := range adj[course] {
			if !dfs(prereq) {
				return false
			}
		}
		visited[course] = false
		done[course] = true
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}
	return true
}
