type Item struct {
	val   int
	count int
}

type MaxHeap []*Item

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	h := &MaxHeap{}
	heap.Init(h)
	for val, count := range counts {
		heap.Push(h, &Item{val: val, count: count})
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(h).(*Item)
		res[i] = item.val
	}

	return res
}
