func rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1
	for l < r {
		for i := 0; i < r-l; i++ {
			top, bottom := l, r

			// save the topleft
			topLeft := matrix[top][l+i]

			// move bottom left into top left
			matrix[top][l+i] = matrix[bottom-i][l]

			// move bottom right into bottom left
			matrix[bottom-i][l] = matrix[bottom][r-i]

			// move top right into bottom right
			matrix[bottom][r-i] = matrix[top+i][r]

			// move top left into top right
			matrix[top+i][r] = topLeft
		}
		r--
		l++
	}
}
