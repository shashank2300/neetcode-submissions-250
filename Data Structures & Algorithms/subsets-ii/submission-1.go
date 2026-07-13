func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    ans := make([][]int, 0)

    var recurse func(i int, cur []int)
    recurse = func(i int, cur []int) {
        ans = append(ans, append([]int{}, cur...))

        for j := i; j < len(nums); j++ {
            if j > i && nums[j] == nums[j-1] {
                continue
            }

            cur = append(cur, nums[j])
            recurse(j+1, cur)
            cur = cur[:len(cur)-1]
        }
    }

    recurse(0, []int{})
    return ans
}
