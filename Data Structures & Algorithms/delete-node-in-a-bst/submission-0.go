/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else {
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }

        cur := root.Right
        for cur.Left != nil {
            cur = cur.Left
        }
        cur.Left = root.Left
        return root.Right
    }

    return root
}
