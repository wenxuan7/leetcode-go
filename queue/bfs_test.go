package queue

import (
	"leetcode-go/assert"
	"testing"
)
import . "leetcode-go/data"

func TestLevelOrder(t *testing.T) {
	data := []*TreeNode{
		{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
	}
	actual := [][][]int{
		{{3}, {9, 20}, {15, 7}},
	}

	for i := range data {
		result := levelOrder(data[i])
		assert.Verify2Arr(t, i, result, actual[i])
	}
}

func TestMinMutation(t *testing.T) {
	startGene := []string{
		"AACCGGTT",
		"AAAAAAAT",
	}
	endGene := []string{
		"AACCGGTA",
		"CCCCCCCC",
	}
	bank := [][]string{
		{"AACCGGTA"},
		{"AAAAAAAC", "AAAAAAAA", "CCCCCCCC"},
	}
	actual := []int{
		1,
		-1,
	}

	for i := range startGene {
		result := minMutation(startGene[i], endGene[i], bank[i])
		assert.Verify(t, i, result, actual[i])
	}
}
