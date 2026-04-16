type MinHeap []int

func (h MinHeap) Len() int { return len(h)}
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MaxHeap []int
func (h MaxHeap) Len() int { return len(h)}
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	left  MaxHeap
	right MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  make(MaxHeap, 0),
		right: make(MinHeap, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.left, num)

	heap.Push(&this.right, heap.Pop(&this.left))

	if len(this.left) < len(this.right) {
		heap.Push(&this.left, heap.Pop(&this.right))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.left) > len(this.right) {
		return float64(this.left[0])
	}
	return float64(this.left[0]+this.right[0]) / 2.0
}
