func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, 0)


	for i, v := range temperatures {
		cur := 0
		for len(stack)>0 && v > stack[len(stack)-1] {
			cur++
			if res[i-cur] != 0 {
				continue
			}
			stack = stack[:len(stack)-1]
			res[i-cur] = cur
		}
		stack = append(stack, v)
		// fmt.Println(stack)
		// fmt.Println(res)
	}
	return res
}
