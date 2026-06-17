func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    res := make([]int, len(num1)+len(num2))
    for i1 := len(num1) - 1; i1 >= 0; i1-- {
        for i2 := len(num2) - 1; i2 >= 0; i2-- {
            pos := len(num1) - 1 - i1 + len(num2) - 1 - i2
            digit := int(num1[i1]-'0') * int(num2[i2]-'0')

            res[pos] += digit
            res[pos+1] += res[pos] / 10
            res[pos] = res[pos] % 10
        }
    }

    var result strings.Builder
    start := len(res) - 1
    for start >= 0 && res[start] == 0 {
        start--
    }
    if start < 0 {
        return "0"
    }

    for i := start; i >= 0; i-- {
        result.WriteString(strconv.Itoa(res[i]))
    }

    return result.String()
}