/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    if root == nil || p == nil || q == nil {
		return nil
	}
	if max(p.Val, q.Val) < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if min(p.Val, q.Val) > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

