func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	ans := make([][]int, 0)
	for i := range nums {
		count[nums[i]]--

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i+1; j < len(nums); j++ {
			count[nums[j]]--
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			target := -(nums[i] + nums[j])
			if count[target] > 0 {
				ans = append(ans, []int{nums[i], nums[j], target})
			}
		}

		for j := i+1; j < len(nums); j++ {
			count[nums[j]]++
		}
		
	}

	return ans
}
