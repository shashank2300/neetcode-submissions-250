/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        root = &TreeNode{Val: val}
        return root
    }
    cur := root
    for cur != nil {
        if val < cur.Val {
            if cur.Left == nil {
                cur.Left = &TreeNode{Val: val}
                return root
            }
            cur = cur.Left
        } else {
            if cur.Right == nil {
                cur.Right = &TreeNode{Val: val}
                return root
            }
            cur = cur.Right
        }
    }

    return root
}
