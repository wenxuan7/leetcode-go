package sort

import . "github.com/leetcode-go/data"

// sortColors
// 75. 颜色分类
// https://leetcode.cn/problems/sort-colors/
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	p0, p2 := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			if nums[i] == 2 {
				nums[i], nums[p2] = nums[p2], nums[i]
			}
		}

		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

// insertionSortList
// 147. 对链表进行插入排序
// https://leetcode.cn/problems/insertion-sort-list/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
