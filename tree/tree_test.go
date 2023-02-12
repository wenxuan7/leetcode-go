package tree

import (
	"leetcode-go/assert"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	data := [][]string{
		{"1", "null", "2", "3"},
		{},
		{"1"},
	}
	actual := [][]int{
		{1, 3, 2},
		{},
		{1},
	}

	for i := range data {
		result := inorderTraversal(assert.Generate2TreeOfPreorder(data[i], 0))
		assert.VerifyArr(t, i, result, actual[i])
	}
}

func TestPreorderTraversal(t *testing.T) {
	data := [][]string{
		{"1", "null", "2", "3"},
		{},
		{"1"},
	}
	actual := [][]int{
		{1, 2, 3},
		{},
		{1},
	}

	for i := range data {
		result := preorderTraversal(assert.Generate2TreeOfPreorder(data[i], 0))
		assert.VerifyArr(t, i, result, actual[i])
	}
}

func TestPostorder(t *testing.T) {
	data := [][]string{
		{"1", "null", "3", "2", "4", "null", "5", "6"},
		{},
		{"1"},
	}
	actual := [][]int{
		{5, 6, 3, 2, 4, 1},
		{},
		{1},
	}

	for i := range data {
		result := postorder(assert.GenerateNTreeOfPreorder(data[i]))
		assert.VerifyArr(t, i, result, actual[i])
	}
}

func TestPreorder(t *testing.T) {
	data := [][]string{
		{"1", "null", "3", "2", "4", "null", "5", "6"},
		{},
		{"1"},
	}
	actual := [][]int{
		{1, 3, 5, 6, 2, 4},
		{},
		{1},
	}

	for i := range data {
		result := preorder(assert.GenerateNTreeOfPreorder(data[i]))
		assert.VerifyArr(t, i, result, actual[i])
	}
}

func TestLevelOrder(t *testing.T) {
	data := [][]string{
		{"1", "null", "3", "2", "4", "null", "5", "6"},
		{},
		{"1"},
	}
	actual := [][][]int{
		{{1}, {3, 2, 4}, {5, 6}},
		{},
		{{1}},
	}

	for i := range data {
		result := levelOrder(assert.GenerateNTreeOfPreorder(data[i]))
		assert.VerifySecArr(t, i, result, actual[i])
	}
}
