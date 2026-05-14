/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil{
		return head
	}

	// 1. Count the total number of nodes
    curr := head
    count := 0
    for curr != nil {
        count++
        curr = curr.Next
    }

    // 2. Setup a dummy node to easily handle the head of the whole list
    dummy := &ListNode{Next: head}
    groupPrev := dummy  // The node right before our current k-group
    curr = head         // The first node of our current k-group

    // 3. Process the list in chunks of k
    for count >= k {
        // Find the start of the *next* group 
        // (This will be passed as the 'next' argument to your helper)
        nextGroupHead := curr
        for i := 0; i < k; i++ {
            nextGroupHead = nextGroupHead.Next
        }

        // Reverse the current group.
        // groupPrev.Next is updated to point to the NEW head of the reversed group.
        groupPrev.Next = reverse(curr, nextGroupHead, k)

        // Step forward for the next iteration:
        // After reversal, 'curr' (which was the old head) is now the tail of this group.
        groupPrev = curr 
        curr = nextGroupHead
        
        // Decrement the count by the k nodes we just processed
        count -= k
    }

    return dummy.Next
}

func reverse(head, next *ListNode, k int) *ListNode {
	prev := next
	for _ = range k {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
