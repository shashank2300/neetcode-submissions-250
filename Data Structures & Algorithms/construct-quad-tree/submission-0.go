/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    return dfs(grid, len(grid), 0, 0)
}

func dfs(grid [][]int, n, r, c int) *Node {
    allSame := true
    for i := 0; i < n && allSame; i++ {
        for j := 0; j < n; j++ {
            if grid[r][c] != grid[r+i][c+j] {
                allSame = false
                break
            }
        }
    }

    if allSame {
        return &Node{Val: grid[r][c] == 1, IsLeaf: true}
    }

    mid := n / 2
    topLeft := dfs(grid, mid, r, c)
    topRight := dfs(grid, mid, r, c+mid)
    bottomLeft := dfs(grid, mid, r+mid, c)
    bottomRight := dfs(grid, mid, r+mid, c+mid)

    return &Node{
        Val:         false,
        IsLeaf:      false,
        TopLeft:     topLeft,
        TopRight:    topRight,
        BottomLeft:  bottomLeft,
        BottomRight: bottomRight,
    }
}