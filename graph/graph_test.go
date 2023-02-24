package graph

import (
	"testing"
)

func TestCloneGraph(t *testing.T) {
	data := []*Node{
		generateGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
	}
	actual := []*Node{
		generateGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
	}

	for i := range data {
		result := cloneGraph(data[i])
		bfsCompareNode(t, i, result, actual[i])
	}
}

// generateGraph 领接列表转换为图
// [[2,4],[1,3],[2,4],[1,3]]
func generateGraph(neighbors [][]int) *Node {
	if len(neighbors) == 0 {
		return nil
	}

	m := make(map[int]*Node, len(neighbors))
	var (
		curr         *Node
		currChildren *Node
		ok           bool
	)

	for i, neighbor := range neighbors {
		if curr, ok = m[i+1]; !ok {
			curr = &Node{Val: i + 1}
			m[i+1] = curr
		}

		nodes := make([]*Node, len(neighbor))
		curr.Neighbors = nodes
		for j, val := range neighbor {
			if currChildren, ok = m[val]; !ok {
				nodes[j] = &Node{Val: val}
				m[val] = nodes[j]
				continue
			}

			m[val] = currChildren
			nodes[j] = currChildren
		}
	}

	return m[1]
}

func bfsCompareNode(t *testing.T, caseIndex int, result, actual *Node) {
	qActual := []*Node{actual}
	qResult := []*Node{result}
	m := make(map[int]bool, 0)
	var currResult, currActual *Node = nil, nil

	for len(qActual) > 0 {
		currResult, currActual = qResult[0], qActual[0]
		qResult, qActual = qResult[1:], qActual[1:]
		if currResult == nil && actual == nil {
			continue
		}

		if currResult == nil {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: nil, actual: %d", caseIndex, currActual.Val)
		}
		if currActual == nil {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %d, actual: nil", caseIndex, currResult.Val)
		}
		if currResult.Val != currActual.Val {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %d, actual: %d", caseIndex, currResult.Val, currActual.Val)
		}
		if m[currActual.Val] {
			continue
		}

		if currResult.Neighbors == nil && currActual.Neighbors == nil {
			continue
		}

		if currResult.Neighbors == nil {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: nil, actual: %v", caseIndex, currActual.Neighbors)
		}
		if currActual.Neighbors == nil {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: nil", caseIndex, currResult.Neighbors)
		}
		if len(currActual.Neighbors) != len(currActual.Neighbors) {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v", caseIndex, currResult.Neighbors, currActual.Neighbors)
		}

		m[currActual.Val] = true
		for i := range currActual.Neighbors {
			qResult = append(qResult, currResult.Neighbors[i])
			qActual = append(qActual, currActual.Neighbors[i])
		}
	}

}
