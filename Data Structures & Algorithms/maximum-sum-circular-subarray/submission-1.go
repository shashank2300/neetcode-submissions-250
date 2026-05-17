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

	if maxSum > 0 {
		return max(maxSum, totalSum - minSum)
	}

	return maxSum
}
