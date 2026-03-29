func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sm, tm := make(map[byte]int, 0), make(map[byte]int)
    for i := range s {
        sm[s[i]]++
        tm[t[i]]++
    }

    for k, v := range sm {
        if tm[k] != v {
            return false
        }
    }
    return true
}
