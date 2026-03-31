/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type MinHeap []*ListNode

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	
	if len(lists) == 0 {
		return nil
	}

	h := &MinHeap{}
	heap.Init(h)

	for _, l := range lists {
		heap.Push(h, l)
	}

	dummy := &ListNode{}
	cur := dummy
	for h.Len() > 0 {
		top := heap.Pop(h).(*ListNode)

		cur.Next = top
		cur = cur.Next

		if top.Next != nil {
			heap.Push(h, top.Next)
		}
	}

	return dummy.Next
}
