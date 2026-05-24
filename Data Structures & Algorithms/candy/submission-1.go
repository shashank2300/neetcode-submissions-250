func candy(ratings []int) int {
	n := len(ratings)
	arr := make([]int, n)

	for i := range arr {
		arr[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			arr[i] = arr[i-1] + 1
		}
	}
	for i := n-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if arr[i+1] + 1 > arr[i] {
				arr[i] = arr[i+1] + 1
			}
		}
	}

	ans := 0
	for _, val := range arr {
		ans += val
	}
	return ans
}
