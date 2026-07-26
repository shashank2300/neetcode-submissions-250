func rangeBitwiseAnd(left int, right int) int {
    i := 0
    for left != right {
        left >>= 1
        right >>= 1
        i++
    }
    return left << i
}