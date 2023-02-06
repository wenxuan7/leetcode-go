package list

import . "leetcode-go/data"

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	tail := reverseList(head.Next)
	next.Next = head
	head.Next = nil
	return tail
}
