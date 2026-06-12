func getOrder(tasks [][]int) []int {
    n := len(tasks)
    indices := make([]int, n)
    for i := 0; i < n; i++ {
        indices[i] = i
    }

    sort.Slice(indices, func(i, j int) bool {
        if tasks[indices[i]][0] != tasks[indices[j]][0] {
            return tasks[indices[i]][0] < tasks[indices[j]][0]
        }
        return indices[i] < indices[j]
    })

    minHeap := &minHeap3{tasks: tasks}
    heap.Init(minHeap)

    result := make([]int, 0, n)
    time := 0
    i := 0

    for minHeap.Len() > 0 || i < n {
        for i < n && tasks[indices[i]][0] <= time {
            heap.Push(minHeap, indices[i])
            i++
        }

        if minHeap.Len() == 0 {
            time = tasks[indices[i]][0]
        } else {
            nextIndex := heap.Pop(minHeap).(int)
            time += tasks[nextIndex][1]
            result = append(result, nextIndex)
        }
    }

    return result
}

type minHeap3 struct {
    tasks [][]int
    data  []int
}

func (h minHeap3) Len() int { return len(h.data) }
func (h minHeap3) Less(i, j int) bool {
    a, b := h.data[i], h.data[j]
    if h.tasks[a][1] != h.tasks[b][1] {
        return h.tasks[a][1] < h.tasks[b][1]
    }
    return a < b
}
func (h minHeap3) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *minHeap3) Push(x interface{}) { h.data = append(h.data, x.(int)) }
func (h *minHeap3) Pop() interface{} {
    old := h.data
    n := len(old)
    x := old[n-1]
    h.data = old[0 : n-1]
    return x
}