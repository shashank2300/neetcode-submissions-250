func stoneGameII(piles []int) int {
    n := len(piles)

    suffixSum := make([]int, n)
    suffixSum[n-1] = piles[n-1]
    for i := n - 2; i >= 0; i-- {
        suffixSum[i] = piles[i] + suffixSum[i+1]
    }

    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for i := n - 1; i >= 0; i-- {
        for M := 1; M <= n; M++ {
            for X := 1; X <= 2*M; X++ {
                if i+X > n {
                    break
                }
                dp[i][M] = max(dp[i][M], suffixSum[i]-dp[i+X][max(M, X)])
            }
        }
    }

    return dp[0][1]
}
