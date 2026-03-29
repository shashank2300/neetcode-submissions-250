/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || left == right {
        return head
    }

    dummy := &ListNode{Next: head}
    pre := dummy

    // 1. Reach the node right before the reversal starts
    for i := 1; i < left; i++ {
        pre = pre.Next
    }

    // 2. Start the reversal
    var prev *ListNode
    cur := pre.Next
    for i := 0; i < right-left+1; i++ {
        nextTemp := cur.Next // Store the actual next node
        cur.Next = prev      // Reverse the link
        prev = cur           // Move prev forward
        cur = nextTemp       // Move cur forward
    }

    // 3. Re-connect the reversed segment to the rest of the list
    // pre.Next still points to the node that is now at the END of the reversed segment
    pre.Next.Next = cur  
    pre.Next = prev

    return dummy.Next
}
