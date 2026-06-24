func reverse(x int) int {
    var rec func(int, int) int
    rec = func(n, rev int) int {
        if n == 0 {
            return rev
        }
        rev = rev*10 + n%10
        return rec(n/10, rev)
    }

    sign := 1
    if x < 0 {
        sign = -1
        x = -x
    }

    reversedNum := rec(x, 0) * sign
    if reversedNum < -(1<<31) || reversedNum > (1<<31)-1 {
        return 0
    }
    return reversedNum
}