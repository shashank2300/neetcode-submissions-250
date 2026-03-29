func islandPerimeter(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    visited := make(map[[2]int]bool)
    directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    bfs := func(r, c int) int {
        queue := [][2]int{{r, c}}
        visited[[2]int{r, c}] = true
        perimeter := 0

        for len(queue) > 0 {
            cell := queue[0]
            queue = queue[1:]
            x, y := cell[0], cell[1]

            for _, dir := range directions {
                nx, ny := x+dir[0], y+dir[1]
                if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] == 0 {
                    perimeter++
                } else if !visited[[2]int{nx, ny}] {
                    visited[[2]int{nx, ny}] = true
                    queue = append(queue, [2]int{nx, ny})
                }
            }
        }
        return perimeter
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 {
                return bfs(i, j)
            }
        }
    }
    return 0
}