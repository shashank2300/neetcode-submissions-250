func transpose(matrix [][]int) [][]int {
    ROWS, COLS := len(matrix), len(matrix[0])
    res := make([][]int, COLS)
    for i := range res {
        res[i] = make([]int, ROWS)
    }

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            res[c][r] = matrix[r][c]
        }
    }

    return res
}