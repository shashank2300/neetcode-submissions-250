/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {

	ans := 0
	var depth func(*TreeNode) int

	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)

		ans = max(ans, left+right)
		return max(left, right)+1
	}
	depth(root)
	return ans
}
