func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	l, r := 0, len(people)-1
	boats := 0
	for l <= r {
		if people[r]+people[l] <= limit {
			l++
		}
		boats++
		r--
		
		// fmt.Println(boats)
	}
	// fmt.Println(boats)
	return boats
}
