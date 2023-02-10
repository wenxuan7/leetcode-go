package list

import (
	. "leetcode-go/data"
	"strconv"
	"strings"
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
		verify(t, i, reverseList(generateList(data[i])), generateList(actual[i]))
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
		verify(t, i, swapPairs(generateList(data[i])), generateList(actual[i]))
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
		result := hasCycle(generateCycleList(data[i], pos[i]))
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
		head := generateCycleList(data[i], pos[i])
		result := detectCycle(head)
		actual := getListNode(head, pos[i])

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
		result := reverseKGroup(generateList(data[i]), k[i])
		verify(t, i, result, generateList(actual[i]))
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
		result := mergeTwoLists(generateList(data1[i]), generateList(data2[i]))
		verify(t, i, result, generateList(actual[i]))
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
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %s, actual: %s",
				caseIndex, listToString(result), listToString(actual))
		}

		verify(t, caseIndex, result.Next, actual.Next)
	} else {
		t.Fatalf("结果与实际不相符, caseIndex: %d, result: %s, actual: %s",
			caseIndex, listToString(result), listToString(actual))
	}
}
