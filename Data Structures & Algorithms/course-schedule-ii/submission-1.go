func findOrder(numCourses int, prerequisites [][]int) []int {
    adj := make([][]int, numCourses)
	incoming := make([]int, numCourses)

	for _, prereq := range prerequisites {
		u, v := prereq[0], prereq[1]
		adj[v] = append(adj[v], u)
		incoming[u]++
	}

	q := make([]int, 0)
	for i := range numCourses {
		if incoming[i] == 0 {
			q = append(q, i)
		}
	}

	ans := make([]int, 0)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		ans = append(ans, cur)
		for _, node := range adj[cur] {
			incoming[node]--
			if incoming[node] == 0 {
				q = append(q, node)
			}
		}
	}

	if len(ans) != numCourses {
		return []int{}
	}

	return ans
}
