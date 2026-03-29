/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    count := 0
    var dfs func(root *TreeNode, m int) 
    dfs = func(root *TreeNode, m int) {
        if root == nil {
            return
        }
        if m <= root.Val {
            count++
        }
        dfs(root.Left, max(m, root.Val))
        dfs(root.Right, max(m, root.Val))
    }

    dfs(root, root.Val)
    return count
}
