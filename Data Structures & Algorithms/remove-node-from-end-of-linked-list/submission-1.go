/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	l, r := dummy, head

	for i := 0; i <n; i++ {
		r = r.Next
	}
	for r != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return dummy.Next
}
