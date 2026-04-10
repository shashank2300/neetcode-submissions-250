func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    }

    // dp2 represents dp[i-2] (no of ways to decode until s[i-2])
    // dp1 represents dp[i-1] (no of ways to decode until s[i-1])
    dp2 := 1 
    dp1 := 1 

    for i := 1; i < len(s); i++ {
        current := 0
        
        // 1. Single digit check: if it's not '0', it's a valid single letter
        if s[i] != '0' {
            current += dp1
        }
        
        // 2. Double digit check: check if s[i-1] and s[i] form a number between 10 and 26
        twoDigit := (s[i-1]-'0')*10 + (s[i]-'0')
        if twoDigit >= 10 && twoDigit <= 26 {
            current += dp2
        }
        
        // Move the window forward
        dp2 = dp1
        dp1 = current
    }

    return dp1
}