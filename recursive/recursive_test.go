package recursive

import (
	"github.com/leetcode-go/assert"
	. "github.com/leetcode-go/data"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	data := []int{
		1,
		2,
		3,
		4,
	}
	actual := []map[string]bool{
		{
			"()": true,
		},
		{
			"(())": true,
			"()()": true,
		},
		{
			"()()()": true,
			"(())()": true,
			"((()))": true,
			"()(())": true,
			"(()())": true,
		},
		{
			"(((())))": true,
			"()()()()": true,
			"(())()()": true,
			"(()()())": true,
			"(())(())": true,
			"((()()))": true,
			"((())())": true,
			"((()))()": true,
			"(()(()))": true,
			"(()())()": true,
			"()((()))": true,
			"()(()())": true,
			"()(())()": true,
			"()()(())": true,
		},
	}

	for i := range data {
		result := generateParenthesis(data[i])
		if len(result) != len(actual[i]) {
			t.Fatalf("结果与实际不相符, resultLen: %d, actualLen: %d", len(result), len(actual[i]))
		}

		for k := range result {
			b, ok := actual[i][result[k]]
			if !ok {
				t.Fatalf("结果与实际不相符, result: %s", result[k])
			}
			if !b {
				t.Fatalf("结果有重复, result: %s", result[k])
			}

			actual[i][result[k]] = false
		}
	}
}

func TestInvertTree(t *testing.T) {
	data := []*TreeNode{
		{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
	}
	actual := []*TreeNode{
		{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
	}

	for i := range data {
		result := invertTree(data[i])
		assert.Verify2Tree(t, i, result, actual[i])
	}
}

func TestMaxDepth(t *testing.T) {
	data := []*TreeNode{
		{Val: 2, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 10, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 17}}},
	}
	actual := []int{
		3,
	}

	for i := range data {
		result := maxDepth(data[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestMinDepth(t *testing.T) {
	data := []*TreeNode{
		{Val: 2, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 10, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 17}}},
	}
	actual := []int{
		2,
	}

	for i := range data {
		result := minDepth(data[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestIsValidBST(t *testing.T) {
	data := []*TreeNode{{
		Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
	}
	actual := []bool{
		true,
	}

	for i := range data {
		result := isValidBST(data[i])
		if result != actual[i] {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v", i, result, actual[i])
		}
	}
}

func TestCodec(t *testing.T) {
	data := []*TreeNode{
		{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}},
	}
	for i := range data {
		codec := Constructor()
		serialize := codec.serialize(data[i])
		result := codec.deserialize(serialize)
		assert.Verify2Tree(t, i, result, data[i])
	}
}
func TestPermute(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
	}
	actual := [][][]int{
		{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
	}

	for i := range data {
		result := permute(data[i])
		assert.Verify2Arr(t, i, result, actual[i])
	}
}

func TestPermuteUnique(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{1, 3, 3},
	}
	actual := [][][]int{
		{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		{{1, 3, 3}, {3, 1, 3}, {3, 3, 1}},
	}

	for i := range data {
		result := permuteUnique(data[i])
		assert.Verify2Arr(t, i, result, actual[i])
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	data := []string{
		"3,5,1,6,2,0,8,nil,nil,7,4",
	}
	p := []*TreeNode{
		{Val: 5},
	}
	q := []*TreeNode{
		{Val: 1},
	}
	actual := []*TreeNode{
		{Val: 3},
	}
	codec := Constructor()
	for i := range data {
		result := lowestCommonAncestor(codec.deserialize(data[i]), p[i], q[i])
		assert.Verify(t, i, result.Val, actual[i].Val)
	}
}

func TestCombine(t *testing.T) {
	data := []int{
		4,
	}
	k := []int{
		2,
	}
	actual := [][][]int{
		{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
	}
	for i := range data {
		result := combine(data[i], k[i])
		assert.Verify2Arr(t, i, result, actual[i])
	}
}

func TestMyPow(t *testing.T) {
	data := []float64{
		2.000,
		2.00000,
		2.10000000,
	}
	data1 := []int{
		10,
		-2,
		3,
	}
	actual := []float64{
		1024.00,
		0.25,
		9.261000000,
	}

	for i := range data {
		result := myPow(data[i], data1[i])
		assert.VerifyFloat64(t, i, result, actual[i], 0.000001)
	}
}

func TestSubsets(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
	}
	actual := [][][]int{
		{{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, {}},
	}

	for i := range data {
		result := subsets(data[i])
		assert.Verify2Arr(t, i, result, actual[i])
	}
}
