func combinationSum(nums []int, target int) [][]int {
	dp := make([][][]int, target+1)

	dp[0] = [][]int{{}}

	for _, num := range nums {
		for i := num; i <= target; i++ {

			if len(dp[i-num]) > 0 {
				for _, prevComb := range dp[i-num] {
					newComb := make([]int, len(prevComb), len(prevComb)+1)
					copy(newComb, prevComb)
					newComb = append(newComb, num)
					
					dp[i] = append(dp[i], newComb)
				}
			}
		}
	}

	return dp[target]
}
