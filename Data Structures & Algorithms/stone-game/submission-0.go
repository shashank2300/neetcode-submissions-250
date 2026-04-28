func stoneGame(piles []int) bool {
    n := len(piles)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    total := 0
    for _, pile := range piles {
        total += pile
    }

    var dfs func(l, r int) int
    dfs = func(l, r int) int {
        if l > r {
            return 0
        }
        if dp[l][r] != -1 {
            return dp[l][r]
        }
        even := (r-l)%2 == 0
        left, right := 0, 0
        if even {
            left = piles[l]
            right = piles[r]
        }
        dp[l][r] = max(dfs(l+1, r)+left, dfs(l, r-1)+right)
        return dp[l][r]
    }

    aliceScore := dfs(0, n-1)
    return aliceScore > total-aliceScore
}