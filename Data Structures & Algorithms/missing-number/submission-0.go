func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return ((n*(n+1))/2) - sum
}
