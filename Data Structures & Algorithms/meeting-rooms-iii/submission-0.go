type RoomHeap [][2]int64

func (h RoomHeap) Len() int { return len(h) }
func (h RoomHeap) Less(i, j int) bool { 
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
	
}
func (h RoomHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *RoomHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int64))
}

func (h *RoomHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	available := &MinHeap{}
    heap.Init(available)
    for i := 0; i < n; i++ {
        heap.Push(available, i)
    }

	used := &RoomHeap{}
    heap.Init(used)
    count := make([]int, n)

    for _, meeting := range meetings {
        start, end := int64(meeting[0]), int64(meeting[1])

		// refresh available rooms heap
		// by checking which meetings ended by start of cur meeting
        for used.Len() > 0 && (*used)[0][0] <= start {
            room := heap.Pop(used).([2]int64)[1]
            heap.Push(available, int(room))
        }

		// if no room is available
		// remove the earliest started meeting
		// and the end of current meeting will be delayed
		// as we have to wait till room becomes available
        if available.Len() == 0 {
            top := heap.Pop(used).([2]int64)
            end = top[0] + (end - start)
            heap.Push(available, int(top[1]))
        }

        room := heap.Pop(available).(int)
        heap.Push(used, [2]int64{end, int64(room)})
        count[room]++
    }

    maxRoom := 0
    for i := 1; i < n; i++ {
        if count[i] > count[maxRoom] {
            maxRoom = i
        }
    }
    return maxRoom
}
