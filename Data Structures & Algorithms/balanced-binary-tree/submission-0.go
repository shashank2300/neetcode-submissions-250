/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Result struct {
	height int
	balanced bool
}

func isBalanced(root *TreeNode) bool {
    var dfs func(*TreeNode) Result
	dfs = func(root *TreeNode) Result {
		if root == nil {
			return Result{0, true}
		}

		left := dfs(root.Left)
		right := dfs(root.Right)
		balanced := left.balanced && right.balanced && abs(left.height-right.height) <= 1
		return Result{max(left.height, right.height)+1, balanced}
	}
	return dfs(root).balanced
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
