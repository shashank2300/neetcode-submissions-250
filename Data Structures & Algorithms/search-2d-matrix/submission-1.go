func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m-1

	for l <= r {
		m := l + (r-l)/2

		if matrix[m][0] > target {
			r = m-1
		} else if matrix[m][n-1] < target {
			l = m+1
		} else {
			l, r := 0, n-1
			for l <= r {
				t := l + (r-l)/2
				if target == matrix[m][t] {
					return true
				} else if target > matrix[m][t] {
					l = t+1
				} else {
					r = t-1
				}
			}
			return false
		}
	}
	return false
}
