func findJudge(n int, trust [][]int) int {
    incoming := make([]int, n+1)
    outgoing := make([]int, n+1)

    for _, t := range trust {
        outgoing[t[0]]++
        incoming[t[1]]++
    }

    for i := 1; i <= n; i++ {
        if outgoing[i] == 0 && incoming[i] == n-1 {
            return i
        }
    }

    return -1
}