package require

import (
	. "github.com/leetcode/data"
	"testing"
)

// HListNodeFromArray 从array生成链表，返回头节点
func HListNodeFromArray(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	preHead := &ListNode{}
	pre := preHead
	var curr *ListNode

	for i := range nums {
		curr = &ListNode{Val: nums[i], Next: nil}
		pre.Next = curr
		pre = curr
	}
	return preHead.Next
}

// EqualListNode 校验两个链表的全部节点值是否相等
func EqualListNode(t *testing.T, caseIndex int, result *ListNode, actual *ListNode) {
	if result == nil && actual == nil {
		return
	}
	str1, str2 := result.String(), actual.String()
	if str1 != str2 {
		t.Fatalf("结果与实际不相符\ncaseIndex: %d\nresult: %s\nactual: %s",
			caseIndex, str1, str2)
	}
}
