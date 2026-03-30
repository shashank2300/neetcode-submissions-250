type Item struct {
	val   int
	count int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	for val, count := range counts {
		heap.Push(pq, &Item{val: val, count: count})
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(pq).(*Item)
		res[i] = item.val
	}

	return res
}
