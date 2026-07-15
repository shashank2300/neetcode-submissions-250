type MaxHeap [][]int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][0] > h[j][0]
	}
	return h[i][1] > h[j][1]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {
	h := &MaxHeap{}
	heap.Init(h)
	if a > 0 {
		heap.Push(h, []int{'a', a})
	}
	if b > 0 {
		heap.Push(h, []int{'b', b})
	}
	if c > 0 {
		heap.Push(h, []int{'c', c})
	}
	

	ans := []byte{}
	for h.Len() > 0 {
		top := heap.Pop(h).([]int)
		count, char := top[1], byte(top[0])

		if len(ans) > 1 && ans[len(ans)-1] == char && ans[len(ans)-2] == char {
			if h.Len() == 0 {
				break
			}
			second := heap.Pop(h).([]int)
			count2, char2 := second[1], byte(second[0])

			ans = append(ans, char2)
			count2--
			if count2 > 0 {
				heap.Push(h, []int{int(char2), count2})
			}
			heap.Push(h, []int{int(char), count})
		} else {
			ans = append(ans, char)
			count--
			if count > 0 {
				heap.Push(h, []int{int(char), count})
			}
		}
	}

	return string(ans)
}
