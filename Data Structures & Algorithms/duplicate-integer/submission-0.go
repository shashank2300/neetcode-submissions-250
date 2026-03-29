func hasDuplicate(nums []int) bool {
    m := make(map[int]bool, 0)
    for _, n := range nums {
        if m[n] {
            return true
        }
        m[n] = true
    }
    return false
}
