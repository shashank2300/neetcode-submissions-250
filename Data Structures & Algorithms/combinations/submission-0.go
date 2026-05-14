func combine(n int, k int) [][]int {
    res := [][]int{}

    var dfs func(int, []int)
    dfs = func(i int, cur []int) {
        if len(cur) == k {
            temp := make([]int, k)
            copy(temp, cur)
            res = append(res, temp)
            return
        }

        for j := i; j <= n; j++ {
            cur = append(cur, j)
            dfs(j+1, cur)
            cur = cur[:len(cur)-1]
        }
    }

    dfs(1, []int{})
    return res
}
