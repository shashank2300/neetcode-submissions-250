type item struct {
    distance float64
    point []int
}

type MinHeap []item

func (h MinHeap) Len() int { return len(h)}
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(item))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func kClosest(points [][]int, k int) [][]int {
    distances := &MinHeap{}
    heap.Init(distances)

    for _, point := range points {
        dist := distance(point[0], point[1])
        heap.Push(distances, item{dist, point})
    }

    ans := [][]int{}
    for _ = range k {
        x := heap.Pop(distances).(item)
        ans = append(ans, x.point)
    }

    return ans
}

func distance(x, y int) float64 {
    return math.Sqrt(float64(x*x + y*y))
}
