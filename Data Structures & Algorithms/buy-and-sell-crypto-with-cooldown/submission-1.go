func maxProfit(prices []int) int {
	n := len(prices)
    memo := make([][2]int, n)
    for i := range memo {
        memo[i][0], memo[i][1] = -1, -1
    }

    var dfs func(i int, buying bool) int
    dfs = func(i int, buying bool) int {
        if i >= n {
            return 0
        }

        bIdx := 0
        if buying {
            bIdx = 1
        }

        if memo[i][bIdx] != -1 {
            return memo[i][bIdx]
        }

        res := dfs(i+1, buying)

        if buying {
            buy := dfs(i+1, false) - prices[i]
            res = max(res, buy)
        } else {
            sell := dfs(i+2, true) + prices[i]
            res = max(res, sell)
        }

        memo[i][bIdx] = res
        return res
    }

    return dfs(0, true)
}