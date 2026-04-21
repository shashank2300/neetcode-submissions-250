func plusOne(digits []int) []int {
	n := len(digits)
    carry := 0
	digits[n-1] += 1
	if digits[n-1] > 9 {
		carry = 1
		digits[n-1] -= 10
	}
	i := 2
	for carry > 0 && i <= n {
		carry = 0
		digits[n-i] += 1
		if digits[n-i] > 9 {
			carry = 1
			digits[n-i] -= 10
		}
		i++
	}
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
