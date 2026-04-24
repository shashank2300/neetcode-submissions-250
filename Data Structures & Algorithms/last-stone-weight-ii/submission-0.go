func lastStoneWeightII(stones []int) int {
    stoneSum := 0
    for _, stone := range stones {
        stoneSum += stone
    }
    target := stoneSum / 2
    n := len(stones)

    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, target+1)
    }

    for i := 1; i <= n; i++ {
        for t := 0; t <= target; t++ {
            if t >= stones[i-1] {
                dp[i][t] = max(dp[i-1][t], dp[i-1][t-stones[i-1]]+stones[i-1])
            } else {
                dp[i][t] = dp[i-1][t]
            }
        }
    }
    return stoneSum - 2*dp[n][target]
}