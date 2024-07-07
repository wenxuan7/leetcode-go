package require

import (
	. "github.com/wenxuan7/leetcode/data"
	"testing"
)

func HListNodeFromArray(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	preH := &ListNode{}
	pre := preH
	for _, v := range nums {
		curr := &ListNode{Val: v}
		pre.Next = curr
		pre = curr
	}
	return preH.Next
}

func EqualListNode(t *testing.T, i int, result, actual *ListNode) {
	resultStr, actualStr := result.String(), actual.String()
	if resultStr != actualStr {
		t.Fatalf("结果错误, case: %d, result: %s, actual: %s",
			i, resultStr, actualStr)
	}
}
