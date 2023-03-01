package assert

import (
	. "github.com/leetcode-go/data"
	"strconv"
	"testing"
)

// Generate2TreeOfPreorder 根据前序遍历生成二叉树
// i 初始位置
func Generate2TreeOfPreorder(nums []string, i int) *TreeNode {
	if len(nums) == 0 || i >= len(nums) ||
		nums[i] == "null" {
		return nil
	}

	var (
		v   int
		err error
	)

	if v, err = strconv.Atoi(nums[i]); err != nil {
		panic(err)
	}

	root := &TreeNode{Val: v}
	root.Left = Generate2TreeOfPreorder(nums, i+1)
	root.Right = Generate2TreeOfPreorder(nums, i+2)
	return root
}

// Verify2Tree 比较二叉树
func Verify2Tree(t *testing.T, caseIndex int, result, actual *TreeNode) {
	if result == nil && actual == nil {
		return
	}

	if result == nil {
		t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
			caseIndex, nil, actual.Val)
	}
	if actual == nil {
		t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
			caseIndex, result.Val, nil)
	}

	if result.Val != actual.Val {
		t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
			caseIndex, result.Val, actual.Val)
	}
	Verify2Tree(t, caseIndex, result.Left, actual.Left)
	Verify2Tree(t, caseIndex, result.Right, actual.Right)
}

// GenerateNTreeOfPreorder 根据层序遍历生成N叉树
func GenerateNTreeOfPreorder(nums []string) *Node {
	if len(nums) == 0 {
		return nil
	}

	var (
		v   int
		err error
	)

	// 队列
	q := make([]*Node, 0, len(nums))
	if v, err = strconv.Atoi(nums[0]); err != nil {
		panic(err)
	}

	root := &Node{Val: v}
	q = append(q, root)
	for j := 2; j < len(nums); j++ {
		var children []*Node
		if j < len(nums) && nums[j] != "null" {
			children = make([]*Node, 0)
		}

		for j < len(nums) && nums[j] != "null" {
			if v, err = strconv.Atoi(nums[j]); err != nil {
				panic(err)
			}

			node := &Node{Val: v}
			children = append(children, node)
			// 入队列
			q = append(q, node)
			j++
		}
		q[0].Children = children
		// 出队列
		q = q[1:]
	}

	return root
}

// VerifyNTree 比较N叉树
func VerifyNTree(t *testing.T, caseIndex int, result, actual *Node) {
	if result == nil && actual == nil {
		return
	}

	if result == nil {
		t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
			caseIndex, nil, actual.Val)
	}
	if actual == nil {
		t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
			caseIndex, result.Val, nil)
	}

	if result.Val != actual.Val {
		t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
			caseIndex, result.Val, actual.Val)
	}
	if len(result.Children) != len(actual.Children) {
		t.Fatalf("结果与实际不相符, caseIndex: %d, resultChildrenLen: %d, actualChildrenLen: %d",
			caseIndex, len(result.Children), len(actual.Children))
	}

	for i := range actual.Children {
		VerifyNTree(t, caseIndex, result.Children[i], actual.Children[i])
	}
}
