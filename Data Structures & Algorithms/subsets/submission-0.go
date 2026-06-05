func subsets(nums []int) [][]int {

    n := len(nums)
    ans := make([][]int, 0)
    cur := []int{}
    var backtrack func(i int)
    backtrack = func(i int) {
        if i >= n {
            temp := make([]int, len(cur))
            copy(temp, cur)
            ans = append(ans, temp)

            return
        }

        cur = append(cur, nums[i])
        backtrack(i+1)
        cur = cur[:len(cur)-1]
        backtrack(i+1)
    }

    backtrack(0)
    return ans
}
