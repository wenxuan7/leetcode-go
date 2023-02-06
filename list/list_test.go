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
		verify(t, i, reverseList(listGenerator(reverseListData[i])), listGenerator(reverseListActual[i]))
	}
}

func listGenerator(num []int) *ListNode {
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

	} else {
		t.Fatal(fmt.Sprintf("结果与实际不相符, caseIndex: %d, result: %s, actual: %s", caseIndex, listToString(result), listToString(actual)))
	}
}
