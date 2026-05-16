func lemonadeChange(bills []int) bool {
	change := make(map[int]int)

	for _, bill := range bills {
		if bill == 5 {
			change[5]++
			continue
		}
		if bill == 10 {
			if change[5] >= 1 {
				change[5]--
				change[10]++
			} else {
				return false
			}
		} else if bill == 20 {
			if change[10] >= 1 && change[5] >= 1 {
				change[10]--
				change[5]--
			} else if change[5] >= 3 {
				change[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
