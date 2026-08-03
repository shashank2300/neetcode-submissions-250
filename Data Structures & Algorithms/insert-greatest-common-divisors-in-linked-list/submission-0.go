/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func gcd(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}
	return gcd(b, a%b)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    cur := head

	for cur != nil {
		if cur.Next == nil {
			break
		}
		next := cur.Next
		cur.Next = &ListNode{gcd(cur.Val, next.Val), next}

		cur = next
	}

	return head
}
