func checkValidString(s string) bool {
	minOpen, maxOpen := 0, 0
	for _, c := range s {
		if c == '(' {
			minOpen++
			maxOpen++
		} else if c == ')' {
			minOpen--
			maxOpen--
		} else {
			minOpen-- // considering * as )
			maxOpen++ // considering * as (
		}

		if maxOpen < 0 {
			return false
		}

		if minOpen < 0 {
			minOpen = 0
		}
	}

	return minOpen == 0
}
