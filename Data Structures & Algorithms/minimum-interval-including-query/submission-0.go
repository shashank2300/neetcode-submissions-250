// 1. Define a custom type for the min-heap
// Each element is a pair: [2]int{interval_size, interval_end}
type minHeap [][2]int

// 2. Implement sort.Interface methods for the heap
func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][0] < h[j][0] } // Min-heap based on size
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 3. Implement heap.Interface methods (Push and Pop must use pointer receivers)
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	// Sort intervals by their start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Pair queries with their original indices and sort them
	queriesWithIdx := make([][2]int, len(queries))
	for i, q := range queries {
		queriesWithIdx[i] = [2]int{q, i}
	}
	sort.Slice(queriesWithIdx, func(i, j int) bool {
		return queriesWithIdx[i][0] < queriesWithIdx[j][0]
	})

	// Initialize the custom min-heap
	h := &minHeap{}
	heap.Init(h)
	
	res := make([]int, len(queries))
	i := 0

	for _, qPair := range queriesWithIdx {
		q, originalIdx := qPair[0], qPair[1]

		// Push all valid intervals into the heap
		for i < len(intervals) && intervals[i][0] <= q {
			size := intervals[i][1] - intervals[i][0] + 1
			heap.Push(h, [2]int{size, intervals[i][1]})
			i++
		}

		// Remove intervals from the heap that end before the current query
		for h.Len() > 0 && (*h)[0][1] < q {
			heap.Pop(h)
		}

		// The top of the heap is the smallest valid interval size
		if h.Len() > 0 {
			res[originalIdx] = (*h)[0][0]
		} else {
			res[originalIdx] = -1
		}
	}

	return res
}