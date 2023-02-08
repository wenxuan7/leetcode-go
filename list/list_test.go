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
	hasCyclePos = []int{
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
		result := hasCycle(generateCycleList(hasCycleData[i], hasCyclePos[i]))
		if result != hasCycleActual[i] {
			t.Fatal(fmt.Sprintf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
				i, result, hasCycleActual[i]))
		}
	}
}

var (
	detectCycleData = [][]int{
		{3, 2, 0, -4},
		{1, 2},
		{},
		{1},
	}
	detectCyclePos = []int{
		1,
		0,
		0,
		-1,
	}
)

func TestDetectCycle(t *testing.T) {
	for i := range detectCycleData {
		head := generateCycleList(detectCycleData[i], detectCyclePos[i])
		result := detectCycle(head)
		actual := getListNode(head, detectCyclePos[i])

		if result != actual {
			t.Fatal(fmt.Sprintf("结果与实际不相符, caseIndex: %d, result: %d, actual: %d",
				i, result.Val, actual.Val))
		}
	}
}

var (
	reverseKGroupData = [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}

	reverseKGroupK = []int{
		2,
		2,
		2,
		2,
		2,
		4,
	}
	reverseKGroupActual = [][]int{
		{},
		{1},
		{2, 1},
		{2, 1, 3},
		{2, 1, 4, 3},
		{4, 3, 2, 1},
	}
)

func TestReverseKGroup(t *testing.T) {
	for i := range reverseKGroupData {
		result := reverseKGroup(generateList(reverseKGroupData[i]), reverseKGroupK[i])
		verify(t, i, result, generateList(reverseKGroupActual[i]))
	}
}

var (
	mergeTwoListsData1 = [][]int{
		{},
		{},
		{1},
		{1, 3, 4, 6},
	}
	mergeTwoListsData2 = [][]int{
		{},
		{1},
		{},
		{2, 5},
	}
	mergeTwoListsActual = [][]int{
		{},
		{1},
		{1},
		{1, 2, 3, 4, 5, 6},
	}
)

func TestMergeTwoLists(t *testing.T) {
	for i := range mergeTwoListsData1 {
		result := mergeTwoLists(generateList(mergeTwoListsData1[i]), generateList(mergeTwoListsData2[i]))
		verify(t, i, result, generateList(mergeTwoListsActual[i]))
	}
}

// getListNode 根据索引获取链表节点
func getListNode(head *ListNode, i int) *ListNode {
	var (
		count = 0
		curr  = head
	)

	for curr != nil {
		if count == i {
			return curr
		}
		curr = curr.Next
		count++
	}

	return nil
}

// generateCycleList 生成环形链表， pos为成环索引位置
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

// generateList 生成链表，返回头节点
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

// listToString 链表转字符串做格式化输出
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

// verify 校验两个链表的全部节点值是否相等
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
