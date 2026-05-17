func maxSubarraySumCircular(nums []int) int {
	totalSum := 0
	maxSum, curMax := nums[0], 0
	minSum, curMin := nums[0], 0
	for _, num := range nums {
		totalSum += num

		curMax = max(curMax + num, num)
		maxSum = max(maxSum, curMax)

		curMin = min(num, curMin + num)
		minSum = min(minSum, curMin)
	}

	// Edge Case: If all numbers are negative, maxSum will be negative.
    // In this case, totalSum == minSum, so totalSum - minSum = 0.
    // But we must pick at least one element, so returning 0 is wrong. 
    // We just return the largest single negative number (maxSum).
	if maxSum > 0 {
		return max(maxSum, totalSum - minSum)
	}

	return maxSum
}
