/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func rob(root *TreeNode) int {
    res := dfs(root)
    return max(res[0], res[1])
}

// dfs returns an array of two integers:
// [0] -> Max money if we DON'T rob the current node
// [1] -> Max money if we DO rob the current node
func dfs(node *TreeNode) [2]int {
    if node == nil {
        return [2]int{0, 0}
    }

    // Post-order traversal: process children first
    left := dfs(node.Left)
    right := dfs(node.Right)

    // Option 1: We DON'T rob this node.
    // So we can choose the max of robbing or not robbing the left child,
    // plus the max of robbing or not robbing the right child.
    notRobbed := max(left[0], left[1]) + max(right[0], right[1])

    // Option 2: We DO rob this node.
    // This means we CANNOT rob its direct children.
    robbed := node.Val + left[0] + right[0]

    return [2]int{notRobbed, robbed}
}
