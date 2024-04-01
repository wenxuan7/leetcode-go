package require

import (
	. "github.com/leetcode/data"
	"strconv"
	"strings"
	"testing"
)

// GetListNode 根据索引获取链表节点
func GetListNode(head *ListNode, i int) *ListNode {
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

// GenerateCycleList 生成环形链表， pos为成环索引位置
func GenerateCycleList(num []int, pos int) *ListNode {
	if len(num) == 0 {
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

// GenerateList 生成链表，返回头节点
func GenerateList(num []int) *ListNode {
	if len(num) == 0 {
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

// ListToString 链表转字符串做格式化输出
func ListToString(node *ListNode) string {
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

// VerifyList 校验两个链表的全部节点值是否相等
func VerifyList(t *testing.T, caseIndex int, result *ListNode, actual *ListNode) {
	if result == nil && actual == nil {
		return
	} else if result != nil && actual != nil {
		if result.Val != actual.Val {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %s, actual: %s",
				caseIndex, ListToString(result), ListToString(actual))
		}

		VerifyList(t, caseIndex, result.Next, actual.Next)
	} else {
		t.Fatalf("结果与实际不相符, caseIndex: %d, result: %s, actual: %s",
			caseIndex, ListToString(result), ListToString(actual))
	}
}
