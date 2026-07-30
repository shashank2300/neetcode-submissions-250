type CountSquares struct {
    ptsCount map[int]map[int]int
}

func Constructor() CountSquares {
    return CountSquares{
        ptsCount: make(map[int]map[int]int),
    }
}

func (this *CountSquares) Add(point []int) {
    x, y := point[0], point[1]
    if this.ptsCount[x] == nil {
        this.ptsCount[x] = make(map[int]int)
    }
    this.ptsCount[x][y]++
}

func (this *CountSquares) Count(point []int) int {
    res := 0
    x1, y1 := point[0], point[1]

    for y2 := range this.ptsCount[x1] {
        side := y2 - y1
        if side == 0 {
            continue
        }

        x3, x4 := x1+side, x1-side

        if _, exists := this.ptsCount[x3]; exists {
            res += this.ptsCount[x1][y2] * this.ptsCount[x3][y1] * this.ptsCount[x3][y2]
        }

        if _, exists := this.ptsCount[x4]; exists {
            res += this.ptsCount[x1][y2] * this.ptsCount[x4][y1] * this.ptsCount[x4][y2]
        }
    }
    return res
}