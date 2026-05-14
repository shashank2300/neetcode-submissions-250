func combine(n int, k int) [][]int {
	var (
		res = make([][]int, 0)
		backtrack func(int, []int)
	)
	backtrack = func(i int, comb []int) {
		if len(comb) == k {
			combCopy := make([]int, len(comb))
			copy(combCopy, comb)
			res = append(res, combCopy)
			return
		}

		if i > n {
			return
		}

		for j := i; j <= n; j++ {
			choice := append(comb, j)
			backtrack(j+1, choice)
		}
	}

	backtrack(1, []int{})

	return res
}
