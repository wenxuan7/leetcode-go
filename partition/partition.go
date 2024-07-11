package partition

import (
	. "github.com/wenxuan7/leetcode/data"
)

// sortedArrayToBST 108. 将有序数组转换为二叉搜索树
// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/?envType=study-plan-v2&envId=top-interview-150
func sortedArrayToBST(nums []int) *TreeNode {
	var partitions func(l, r int) *TreeNode
	partitions = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := l + (r-l)>>2
		node := &TreeNode{Val: nums[mid]}
		node.Left = partitions(l, mid-1)
		node.Right = partitions(mid+1, r)
		return node
	}
	return partitions(0, len(nums)-1)
}

// mergeKLists 23. 合并 K 个升序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/submissions/545898269/?envType=study-plan-v2&envId=top-interview-150
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}

	var recursiveMerge func(l, r int) *ListNode
	recursiveMerge = func(l, r int) *ListNode {
		if l > r {
			return nil
		}
		if l == r {
			return lists[l]
		}
		mid := l + (r-l)>>2
		return merge(recursiveMerge(l, mid), recursiveMerge(mid+1, r))
	}
	return recursiveMerge(0, l-1)
}

func merge(l, r *ListNode) *ListNode {
	if l == r {
		return l
	}
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}

	dummy := &ListNode{Val: -1}
	cur, cur1, cur2 := dummy, l, r
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	for cur1 != nil {
		cur.Next = cur1
		cur1 = cur1.Next
		cur = cur.Next
	}
	for cur2 != nil {
		cur.Next = cur2
		cur2 = cur2.Next
		cur = cur.Next
	}
	return dummy.Next
}
