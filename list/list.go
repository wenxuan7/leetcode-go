package list

import (
	"container/list"
	. "github.com/leetcode-go/data"
)

// MyLinkedList
// 707. 设计链表
// https://leetcode.cn/problems/design-linked-list/
type MyLinkedList struct {
	head, tail *DNode
	len        int
}
type DNode struct {
	prev, next *DNode
	Val        int
}

func NewMyLinkedList() MyLinkedList {
	head, tail := DNode{}, DNode{}
	head.next, tail.prev = &tail, &head
	return MyLinkedList{head: &head, tail: &tail}
}

func (mll *MyLinkedList) Get(index int) int {
	for curr, i := mll.head.next, 0; curr != mll.tail && i < mll.len; curr, i = curr.next, i+1 {
		if i == index {
			return curr.Val
		}
	}

	return -1
}

func (mll *MyLinkedList) AddAtHead(val int) {
	newDNode := &DNode{Val: val}
	mll.head.next.prev = newDNode
	newDNode.next = mll.head.next
	mll.head.next = newDNode
	newDNode.prev = mll.head
	mll.len += 1
}

func (mll *MyLinkedList) AddAtTail(val int) {
	newDNode := &DNode{Val: val}
	mll.tail.prev.next = newDNode
	newDNode.prev = mll.tail.prev
	mll.tail.prev = newDNode
	newDNode.next = mll.tail
	mll.len += 1
}

func (mll *MyLinkedList) AddAtIndex(index int, val int) {
	for curr, i := mll.head, -1; curr != mll.tail && i < mll.len; curr, i = curr.next, i+1 {
		if i == index-1 {
			pre, next := curr, curr.next
			newNode := &DNode{Val: val, prev: pre, next: next}
			pre.next = newNode
			next.prev = newNode
			mll.len += 1
			break
		}
	}
}

func (mll *MyLinkedList) DeleteAtIndex(index int) {
	for curr, i := mll.head.next, 0; curr != mll.tail && i < mll.len; curr, i = curr.next, i+1 {
		if i == index {
			pre, next := curr.prev, curr.next
			pre.next = next
			next.prev = pre
			curr.next = nil
			curr.prev = nil
			mll.len -= 1
			break
		}
	}
}

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
