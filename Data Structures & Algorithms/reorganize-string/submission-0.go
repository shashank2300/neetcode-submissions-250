type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func reorganizeString(s string) string {
    freq := make([]int, 26)
    for _, c := range s {
        freq[c-'a']++
    }

    h := &MaxHeap{}
    heap.Init(h)
    for i := 0; i < 26; i++ {
        if freq[i] > 0 {
            heap.Push(h, []int{freq[i], i})
        }
    }

    res := []byte{}
    var prev []int

    for h.Len() > 0 || prev != nil {
        if prev != nil && h.Len() == 0 {
            return ""
        }

        curr := heap.Pop(h).([]int)
        res = append(res, byte(curr[1]+'a'))
        curr[0]--

        if prev != nil {
            heap.Push(h, prev)
            prev = nil
        }

        if curr[0] > 0 {
            prev = curr
        }
    }

    return string(res)
}