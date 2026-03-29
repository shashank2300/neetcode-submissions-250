/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    ans := -1
    rank := 0
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        rank++
        if rank == k {
            ans = root.Val
            return
        }
        dfs(root.Right)
    }

    dfs(root)

    return ans
}
