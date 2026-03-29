func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
    res := [][]int{}

    for i := 0; i < len(nums); i++ {
        a := nums[i]
        if i > 0 && a == nums[i-1] {
            continue
        }
		for j := i+1; j < len(nums); j++ {
			b := nums[j]
			if j > i+1 && b == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
        	for l < r {
            	fourSum := a + b + nums[l] + nums[r]
            	if fourSum > target {
             	   r--
            	} else if fourSum < target {
                	l++
            	} else {
                	res = append(res, []int{a, b, nums[l], nums[r]})
                	l++
                	r--
                	for l < r && nums[l] == nums[l-1] {
                    	l++
                	}
            	}
        	}
		}
    }

    return res
}