func openLock(deadends []string, target string) int {
    visited := make(map[string]bool)
    for _, d := range deadends {
        visited[d] = true
    }

    if visited["0000"] {
        return -1
    }

    // Queue stores pairs of [state, steps]
    type node struct {
        state string
        steps int
    }
    queue := []node{{"0000", 0}}
    visited["0000"] = true

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]

        if curr.state == target {
            return curr.steps
        }

        for i := 0; i < 4; i++ {
            c := curr.state[i]

            // Turn UP
            var up byte
            if c == '9' { up = '0' } else { up = c + 1 }
            nextUp := curr.state[:i] + string(up) + curr.state[i+1:]
            
            if !visited[nextUp] {
                visited[nextUp] = true
                queue = append(queue, node{nextUp, curr.steps + 1})
            }

            // Turn DOWN
            var down byte
            if c == '0' { down = '9' } else { down = c - 1 }
            nextDown := curr.state[:i] + string(down) + curr.state[i+1:]
            
            if !visited[nextDown] {
                visited[nextDown] = true
                queue = append(queue, node{nextDown, curr.steps + 1})
            }
        }
    }

    return -1
}