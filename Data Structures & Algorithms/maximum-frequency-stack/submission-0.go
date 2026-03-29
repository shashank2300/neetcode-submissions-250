type Item struct {
	freq int
	index int
	val int
}

type MaxHeap []Item

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { 
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	}
	return h[i].index > h[j].index
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


type FreqStack struct {
	h *MaxHeap
	index int
	count map[int]int
}

func Constructor() FreqStack {
	h := &MaxHeap{}
	heap.Init(h)
	return FreqStack{
		count: make(map[int]int),
		h: h,
		index: 0,
	}
}

func (this *FreqStack) Push(val int) {
	this.count[val]++
	heap.Push(this.h, Item{freq: this.count[val], index: this.index, val: val})
	this.index++
}

func (this *FreqStack) Pop() int {
	item := heap.Pop(this.h).(Item)
	this.count[item.val]--
	return item.val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */
 