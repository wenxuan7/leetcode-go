package graph

// Node 图节点
type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph
// 133. 克隆图
// https://leetcode.cn/problems/clone-graph/
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	m := make(map[int]*Node, 0)

	root := &Node{Val: node.Val}
	var dfs func(node, newNode *Node)
	dfs = func(node, newNode *Node) {
		if m[node.Val] != nil {
			return
		}

		m[node.Val] = newNode
		if node.Neighbors == nil {
			newNode.Neighbors = nil
			return
		}
		newNeighbors := make([]*Node, len(node.Neighbors))
		for i, neighbor := range node.Neighbors {
			if m[neighbor.Val] != nil {
				newNeighbors[i] = m[neighbor.Val]
				continue
			}

			newNeighbor := &Node{Val: neighbor.Val}
			newNeighbors[i] = newNeighbor
			dfs(neighbor, newNeighbor)
		}
		newNode.Neighbors = newNeighbors
	}

	dfs(node, root)

	return root
}
