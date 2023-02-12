package list

import (
	"leetcode-go/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	var (
		data = [][]int{
			{1, 2, 3, 4, 5, 6, 7},
			{3, 4, 5, 6, 7, 8},
		}
		actual = [][]int{
			{7, 6, 5, 4, 3, 2, 1},
			{8, 7, 6, 5, 4, 3},
		}
	)

	for i := range data {
		assert.VerifyList(t, i, reverseList(assert.GenerateList(data[i])), assert.GenerateList(actual[i]))
	}
}

func TestSwapPairs(t *testing.T) {
	var (
		data = [][]int{
			{1, 2, 3, 4},
			{},
			{1, 2, 3},
		}
		actual = [][]int{
			{2, 1, 4, 3},
			{},
			{2, 1, 3},
		}
	)

	for i := range data {
		assert.VerifyList(t, i, swapPairs(assert.GenerateList(data[i])), assert.GenerateList(actual[i]))
	}
}

func TestHasCycle(t *testing.T) {
	var (
		data = [][]int{
			{3, 2, 0, -4},
			{1, 2},
			{3, 2, 0, -4},
		}
		pos = []int{
			1,
			0,
			-1,
		}
		actual = []bool{
			true,
			true,
			false,
			false,
		}
	)

	for i := range data {
		result := hasCycle(assert.GenerateCycleList(data[i], pos[i]))
		if result != actual[i] {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
				i, result, actual[i])
		}
	}
}

func TestDetectCycle(t *testing.T) {
	var (
		data = [][]int{
			{3, 2, 0, -4},
			{1, 2},
			{},
			{1},
		}
		pos = []int{
			1,
			0,
			0,
			-1,
		}
	)

	for i := range data {
		head := assert.GenerateCycleList(data[i], pos[i])
		result := detectCycle(head)
		actual := assert.GetListNode(head, pos[i])

		if result != actual {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %d, actual: %d",
				i, result.Val, actual.Val)
		}
	}
}

func TestReverseKGroup(t *testing.T) {
	var (
		data = [][]int{
			{},
			{1},
			{1, 2},
			{1, 2, 3},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
		}

		k = []int{
			2,
			2,
			2,
			2,
			2,
			4,
		}
		actual = [][]int{
			{},
			{1},
			{2, 1},
			{2, 1, 3},
			{2, 1, 4, 3},
			{4, 3, 2, 1},
		}
	)

	for i := range data {
		result := reverseKGroup(assert.GenerateList(data[i]), k[i])
		assert.VerifyList(t, i, result, assert.GenerateList(actual[i]))
	}
}

func TestMergeTwoLists(t *testing.T) {
	var (
		data1 = [][]int{
			{},
			{},
			{1},
			{1, 3, 4, 6},
		}
		data2 = [][]int{
			{},
			{1},
			{},
			{2, 5},
		}
		actual = [][]int{
			{},
			{1},
			{1},
			{1, 2, 3, 4, 5, 6},
		}
	)

	for i := range data1 {
		result := mergeTwoLists(assert.GenerateList(data1[i]), assert.GenerateList(data2[i]))
		assert.VerifyList(t, i, result, assert.GenerateList(actual[i]))
	}
}
