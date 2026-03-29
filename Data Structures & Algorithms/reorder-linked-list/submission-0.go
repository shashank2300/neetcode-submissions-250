/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
    if head == nil {
        return
    }

    var rec func(root, cur *ListNode) *ListNode
    rec = func(root, cur *ListNode) *ListNode {
        if cur == nil {
            return root
        }

        root = rec(root, cur.Next)
        if root == nil {
            return nil
        }

        var tmp *ListNode
        if root == cur || root.Next == cur {
            cur.Next = nil
        } else {
            tmp = root.Next
            root.Next = cur
            cur.Next = tmp
        }

        return tmp
    }

    rec(head, head.Next)
}