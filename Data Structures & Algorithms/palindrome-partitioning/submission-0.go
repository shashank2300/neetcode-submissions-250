func partition(s string) [][]string {
    n := len(s)
    ans := make([][]string, 0)
    var back func(int, *[]string)
    back = func(idx int, cur *[]string) {
        if idx >= n {
            partition := make([]string, len(*cur))
            copy(partition, *cur)
            ans = append(ans, partition)
            return
        }

        for j:=idx; j < n; j++ {
            if isPalindrome(s[idx:j+1]) {
                *cur = append(*cur, s[idx:j+1])
                back(j+1, cur)
                *cur = (*cur)[:len(*cur)-1]
            }
        }
    }
    c := []string{}
    back(0, &c)
    return ans
}

func isPalindrome(s string) bool {
    l := len(s)-1
    for i:=0; i<l; i, l = i+1, l-1 {
        if s[i] != s[l] {
            return false
        }
    }
    return true
}
