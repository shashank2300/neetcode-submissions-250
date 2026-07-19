func generateParenthesis(n int) []string {
    ans := make([]string, 0)
    var backtrack func(open int, close int, cur []byte)
    backtrack = func(open int, close int, cur []byte) {
        if len(cur) == n*2 {
            ans = append(ans, string(append([]byte{}, cur...)))
            return
        }
        if open < n {
            cur = append(cur, '(')
            backtrack(open+1, close, cur)
            cur = cur[:len(cur)-1]
        }
        if open > close {
            cur = append(cur, ')')
            backtrack(open, close+1, cur)
            cur = cur[:len(cur)-1]
        }
        
    }

    backtrack(0, 0, []byte{})
    return ans
}
