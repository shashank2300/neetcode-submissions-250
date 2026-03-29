func getConcatenation(nums []int) []int {
    ans := make([]int, len(nums)*2)
    copy(ans, append(nums, nums...))

    return ans
}
