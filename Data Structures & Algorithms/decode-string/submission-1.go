func decodeString(s string) string {
    i := 0
    return helper(s, &i)
}

func helper(s string, i *int) string {
    res := ""
    k := 0

    for *i < len(s) {
        char := s[*i]

        if char >= '0' && char <= '9' {
            k = k*10 + int(char-'0')
        } else if char == '[' {
            *i++ // Move past '['
            innerStr := helper(s, i) // The "New Worker" starts
            for j := 0; j < k; j++ {
                res += innerStr
            }
            k = 0 // Reset multiplier for the next sibling
        } else if char == ']' {
            // No need to increment i here; the caller increments after return
            return res
        } else {
            res += string(char)
        }
        *i++
    }
    return res
}