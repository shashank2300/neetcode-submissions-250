func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	counts := make(map[int]int)
	for _, num := range hand {
		counts[num]++
	}

	for _, num := range hand {
		if counts[num] > 0 {
			for j := num; j < num+groupSize; j++ {
				if counts[j] <= 0 {
					return false
				}
				counts[j]--
			}
		}
	}

	return true
}
