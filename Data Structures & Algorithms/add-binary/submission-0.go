func addBinary(a string, b string) string {
    res := []byte{}
    carry := 0

    aBytes := []byte(a)
    bBytes := []byte(b)
    for i, j := 0, len(aBytes)-1; i < j; i, j = i+1, j-1 {
        aBytes[i], aBytes[j] = aBytes[j], aBytes[i]
    }
    for i, j := 0, len(bBytes)-1; i < j; i, j = i+1, j-1 {
        bBytes[i], bBytes[j] = bBytes[j], bBytes[i]
    }

    n := len(aBytes)
    if len(bBytes) > n {
        n = len(bBytes)
    }

    for i := 0; i < n; i++ {
        digitA := 0
        digitB := 0
        if i < len(aBytes) {
            digitA = int(aBytes[i] - '0')
        }
        if i < len(bBytes) {
            digitB = int(bBytes[i] - '0')
        }

        total := digitA + digitB + carry
        res = append(res, byte(total%2)+'0')
        carry = total / 2
    }

    if carry > 0 {
        res = append(res, '1')
    }

    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
    return string(res)
}