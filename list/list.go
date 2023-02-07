package list

import . "leetcode-go/data"

// reverseList
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

// swapPairs
// 24. 两两交换链表中的节点
// https://leetcode.cn/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	sec := swapPairs(next.Next)
	head.Next = sec
	next.Next = head
	return next
}
