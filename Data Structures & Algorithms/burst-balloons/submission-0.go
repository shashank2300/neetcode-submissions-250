func maxCoins(nums []int) int {
    n := len(nums)
    newNums := append([]int{1}, nums...)
    newNums = append(newNums, 1)

    dp := make([][]int, n+2)
    for i := range dp {
        dp[i] = make([]int, n+2)
    }

    for l := n; l >= 1; l-- {
        for r := l; r <= n; r++ {
            for i := l; i <= r; i++ {
                coins := newNums[l-1] * newNums[i] * newNums[r+1]
                coins += dp[l][i-1] + dp[i+1][r]
                dp[l][r] = max(dp[l][r], coins)
            }
        }
    }

    return dp[1][n]
}
