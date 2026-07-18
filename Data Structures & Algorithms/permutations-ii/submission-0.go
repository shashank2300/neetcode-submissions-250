func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)

    ans := make([][]int, 0)
    var backtrack func([]bool, []int)
    backtrack= func(used []bool, cur []int) {
        if len(cur) == len(nums) {
            ans = append(ans, append([]int{}, cur...))
            return
        }

        for i := range (nums) {
            if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
                continue
            }
            used[i] = true
            cur = append(cur, nums[i])
            backtrack(used, cur)

            used[i] = false
            cur = cur[:len(cur)-1]
        }
    }
    used := make([]bool, len(nums))
    backtrack(used, []int{})

    return ans
}
