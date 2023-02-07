package list

import (
	"fmt"
	. "leetcode-go/data"
	"strconv"
	"strings"
	"testing"
)

var (
	reverseListData = [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8},
	}
	reverseListActual = [][]int{
		{7, 6, 5, 4, 3, 2, 1},
		{8, 7, 6, 5, 4, 3},
	}
)

func TestReverseList(t *testing.T) {
	for i := range reverseListData {
		verify(t, i, reverseList(generateList(reverseListData[i])), generateList(reverseListActual[i]))
	}
}

var (
	swapPairsData = [][]int{
		{1, 2, 3, 4},
		{},
		{1, 2, 3},
	}
	swapPairsActual = [][]int{
		{2, 1, 4, 3},
		{},
		{2, 1, 3},
	}
)

func TestSwapPairs(t *testing.T) {
	for i := range swapPairsData {
		verify(t, i, swapPairs(generateList(swapPairsData[i])), generateList(swapPairsActual[i]))
	}
}

var (
	hasCycleData = [][]int{
		{3, 2, 0, -4},
		{1, 2},
		{3, 2, 0, -4},
	}
	pos = []int{
		1,
		0,
		-1,
	}
	hasCycleActual = []bool{
		true,
		true,
		false,
		false,
	}
)

func TestHasCycle(t *testing.T) {
	for i := range hasCycleData {
		result := hasCycle(generateCycleList(hasCycleData[i], pos[i]))
		if result != hasCycleActual[i] {
			t.Fatal(fmt.Sprintf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
				i, result, hasCycleActual[i]))
		}
	}
}

func generateCycleList(num []int, pos int) *ListNode {
	if num == nil || len(num) == 0 {
		return nil
	}

	var (
		preHead = &ListNode{}
		pre     = preHead
		curr    *ListNode
		posNode *ListNode
		tail    *ListNode
	)

	for i := range num {
		curr = &ListNode{Val: num[i], Next: nil}
		pre.Next = curr
		pre = curr

		if i == pos {
			posNode = curr
		}
		if i == len(num)-1 {
			tail = curr
		}
	}

	if posNode != nil {
		tail.Next = posNode
	}

	return preHead.Next
}

func generateList(num []int) *ListNode {
	if num == nil || len(num) == 0 {
		return nil
	}

	var (
		preHead           = &ListNode{}
		pre               = preHead
		curr    *ListNode = nil
	)

	for i := range num {
		curr = &ListNode{Val: num[i], Next: nil}
		pre.Next = curr
		pre = curr
	}

	return preHead.Next
}

func listToString(node *ListNode) string {
	bd := strings.Builder{}
	bd.WriteString("[")
	for node != nil {
		bd.WriteString(strconv.Itoa(node.Val))
		if node.Next != nil {
			bd.WriteString(", ")
		}
		node = node.Next
	}
	bd.WriteString("]")

	return bd.String()
}

func verify(t *testing.T, caseIndex int, result *ListNode, actual *ListNode) {
	if result == nil && actual == nil {
		return
	} else if result != nil && actual != nil {
		if result.Val != actual.Val {
			t.Fatal(fmt.Sprintf("结果与实际不相符, caseIndex: %d, result: %s, actual: %s",
				caseIndex, listToString(result), listToString(actual)))
		}

		verify(t, caseIndex, result.Next, actual.Next)
	} else {
		t.Fatal(fmt.Sprintf("结果与实际不相符, caseIndex: %d, result: %s, actual: %s",
			caseIndex, listToString(result), listToString(actual)))
	}
}
