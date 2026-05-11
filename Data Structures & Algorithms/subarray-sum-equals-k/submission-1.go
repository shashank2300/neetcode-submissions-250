func subarraySum(nums []int, k int) int {
	count, curSum := 0, 0
	prefixSums := map[int]int{0: 1}

	for _, num := range nums {
		curSum += num
		diff := curSum - k
		count += prefixSums[diff]
		prefixSums[curSum]++
	}

	return count
}
