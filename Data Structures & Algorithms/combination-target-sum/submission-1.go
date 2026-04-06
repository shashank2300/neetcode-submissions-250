func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(nums)

    var dfs func(int, []int, int)
    dfs = func(i int, cur []int, total int) {
        if total == target {
            temp := make([]int, len(cur))
            copy(temp, cur)
            res = append(res, temp)
            return
        }

        for j := i; j < len(nums); j++ {
            if total + nums[j] > target {
                return
            }
            cur = append(cur, nums[j])
            dfs(j, cur, total + nums[j])
            cur = cur[:len(cur)-1]
        }
    }

    dfs(0, []int{}, 0)
    return res
}