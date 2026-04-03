/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
		return nil
	}

	q := make([]*Node, 0)
	q = append(q, node)

	oldToNew := make(map[*Node]*Node, 0)
	oldToNew[node] = &Node{node.Val, make([]*Node, 0)}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, neighbor := range cur.Neighbors {
			if _, ok := oldToNew[neighbor]; !ok {
				oldToNew[neighbor] = &Node{Val: neighbor.Val, Neighbors: make([]*Node, 0)}
				q = append(q, neighbor)
			}
			oldToNew[cur].Neighbors = append(oldToNew[cur].Neighbors, oldToNew[neighbor])
		}
	}
	return oldToNew[node]
}
