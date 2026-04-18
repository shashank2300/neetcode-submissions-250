func setZeroes(matrix [][]int) {
    rows, cols := make(map[int]bool), make(map[int]bool)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := range rows {
		for j := range matrix[0] {
			matrix[i][j] = 0
		}
	}
	for j := range cols {
		for i := range matrix {
			matrix[i][j] = 0
		}
	}
}
