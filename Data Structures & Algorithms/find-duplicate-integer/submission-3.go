func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		// collision point in the cycle
		if slow == fast {
			break
		}
	}

	// distance from start to cycle start
	// is equal to distance from cycle start
	// to collision point
	// this can be proven mathematically
	slow2 := nums[0]
	for slow != slow2{
		slow = nums[slow]
		slow2 = nums[slow2]
	}

	return slow

}
