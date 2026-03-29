type Item struct {
	val int
	idx int
}

type MaxHeap []Item

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	h := &MaxHeap{}
	heap.Init(h)
	ans := make([]int, 0, n-k+1)

	for i := 0; i < n; i++ {
		heap.Push(h, Item{val: nums[i], idx: i})

		for (*h)[0].idx <= i-k {
			heap.Pop(h)
		}

		if i >= k-1 {
			ans = append(ans, (*h)[0].val)
		}
	}

	return ans
}