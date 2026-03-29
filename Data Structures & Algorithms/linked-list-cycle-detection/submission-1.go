/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head== nil || head.Next == nil {
		return false
	}
    slow, fast := head, head.Next

	for slow!=nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		fmt.Println(slow.Val, fast.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
