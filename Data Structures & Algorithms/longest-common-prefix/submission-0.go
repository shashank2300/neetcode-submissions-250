func longestCommonPrefix(strs []string) string {
    sort.Strings(strs)
    a, b := strs[0], strs[len(strs)-1]
    aLen, bLen := len(a), len(b)

    l := 0
    for ; l < aLen && l < bLen && a[l] == b[l]; l++ {

    }

    return a[:l]
}
