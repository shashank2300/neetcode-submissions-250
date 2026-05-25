func combinationSum4(nums []int, target int) int {
    dp := make(map[int]int)
    dp[0] = 1

    for total := 1; total <= target; total++ {
        dp[total] = 0
        for _, num := range nums {
            if val, ok := dp[total-num]; ok {
                dp[total] += val
            }
        }
    }

    return dp[target]
}