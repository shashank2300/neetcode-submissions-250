/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}

	carry := 0

	cur := dummy
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		if sum > 9 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		cur.Next = &ListNode{sum, nil}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + carry
		if sum > 9 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		cur.Next = &ListNode{sum, nil}
		cur = cur.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + carry
		if sum > 9 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		cur.Next = &ListNode{sum, nil}
		cur = cur.Next
		l2 = l2.Next
	}

	if carry > 0 {
		cur.Next = &ListNode{1, nil}
	}
	return dummy.Next
}
