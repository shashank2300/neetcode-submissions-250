func combinationSum2(nums []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(nums)

    var dfs func(i int, cur []int, total int)
    dfs = func(i int, cur []int, total int) {
        // Base cases
        if total == target {
            temp := make([]int, len(cur))
            copy(temp, cur)
            res = append(res, temp)
            return
        }
        if i >= len(nums) || total > target {
            return
        }

        // --- BRANCH 1: INCLUDE the current element ---
        cur = append(cur, nums[i])
        dfs(i+1, cur, total+nums[i])
        cur = cur[:len(cur)-1] // backtrack

        // --- BRANCH 2: SKIP the current element ---
        // To prevent duplicates, if we skip this number, we must 
        // skip all identical numbers that come immediately after it.
        nextIndex := i + 1
        for nextIndex < len(nums) && nums[nextIndex] == nums[i] {
            nextIndex++
        }
        
        dfs(nextIndex, cur, total)
    }

    dfs(0, []int{}, 0)
    return res
}