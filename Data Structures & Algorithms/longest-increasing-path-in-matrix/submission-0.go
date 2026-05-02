func longestIncreasingPath(matrix [][]int) int {
    rows, cols := len(matrix), len(matrix[0])
    indegree := make([][]int, rows)
    for i := range indegree {
        indegree[i] = make([]int, cols)
    }

    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            for _, d := range directions {
                nr, nc := r + d[0], c + d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
                   matrix[nr][nc] < matrix[r][c] {
                    indegree[r][c]++
                }
            }
        }
    }

    queue := [][]int{}
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if indegree[r][c] == 0 {
                queue = append(queue, []int{r, c})
            }
        }
    }

    lis := 0
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            r, c := node[0], node[1]
            for _, d := range directions {
                nr, nc := r + d[0], c + d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
                   matrix[nr][nc] > matrix[r][c] {
                    indegree[nr][nc]--
                    if indegree[nr][nc] == 0 {
                        queue = append(queue, []int{nr, nc})
                    }
                }
            }
        }
        lis++
    }

    return lis
}