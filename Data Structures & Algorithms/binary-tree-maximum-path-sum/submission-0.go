/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    res := root.Val

    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftMax := dfs(node.Left)
        rightMax := dfs(node.Right)

        leftMax = max(leftMax, 0)
        rightMax = max(rightMax, 0)

		// best path including current node
        res = max(res, node.Val+leftMax+rightMax)

		// best downward path from current node
        return node.Val + max(leftMax, rightMax)
    }

    dfs(root)
    return res
}
