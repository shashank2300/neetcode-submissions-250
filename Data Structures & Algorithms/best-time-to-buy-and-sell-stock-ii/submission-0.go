func maxProfit(prices []int) int {
    n := len(prices)
    dp := make([][2]int, n+1)

    for i := n - 1; i >= 0; i-- {
        dp[i][0] = max(dp[i+1][0], -prices[i]+dp[i+1][1])
        dp[i][1] = max(dp[i+1][1], prices[i]+dp[i+1][0])
    }

    return dp[0][0]
}
