package list

import (
	"container/list"
	. "leetcode-go/data"
)

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

// hasCycle
// 141. 环形链表
// https://leetcode.cn/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// detectCycle
// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	fast, slow, hasCycle := head, head, false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

	slow = head
	for fast != nil {
		if slow == fast {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

// reverseKGroup
// 25. K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}

	ls, curr := list.New(), head
	for i := 0; i < k; i++ {
		if curr == nil {
			return head
		}

		ls.PushBack(curr)
		curr = curr.Next
	}

	groupNext := reverseKGroup(curr, k)
	pre := groupNext

	for ls.Len() > 0 {
		element := ls.Front()
		front := element.Value.(*ListNode)
		front.Next = pre
		pre = front
		ls.Remove(element)
	}

	return pre
}

// mergeTwoLists
// 21. 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		next := mergeTwoLists(list1, list2.Next)
		list2.Next = next
		return list2
	} else {
		next := mergeTwoLists(list1.Next, list2)
		list1.Next = next
		return list1
	}
}
