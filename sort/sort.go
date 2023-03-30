package sort

import (
	. "github.com/leetcode-go/data"
	"github.com/leetcode-go/tool"
)

// qSort
// 快排
func qSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	flag, j := l, l+1
	for i := j; i <= r; i++ {
		if nums[i] < nums[flag] {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	// 与j 的前一个交换 j有可能为r+1
	nums[flag], nums[j-1] = nums[j-1], nums[flag]
	flag = j - 1

	qSort(nums, l, flag-1)
	qSort(nums, flag+1, r)
}

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

// maxWidthOfVerticalArea
// 1637. 两点之间不包含任何点的最宽垂直区域
// https://leetcode.cn/problems/widest-vertical-area-between-two-points-containing-no-points/
func maxWidthOfVerticalArea(points [][]int) int {
	if len(points) < 2 {
		return 0
	}

	var qSort func(nums [][]int, l, r int)
	qSort = func(nums [][]int, l, r int) {
		if l >= r {
			return
		}

		flag, j := l, l+1
		for i := j; i <= r; i++ {
			if nums[i][0] < nums[flag][0] {
				nums[j][0], nums[i][0] = nums[i][0], nums[j][0]
				nums[j][1], nums[i][1] = nums[i][1], nums[j][1]
				j++
			}
		}

		nums[flag][0], nums[j-1][0] = nums[j-1][0], nums[flag][0]
		nums[flag][1], nums[j-1][1] = nums[j-1][1], nums[flag][1]
		flag = j - 1

		qSort(nums, l, flag-1)
		qSort(nums, flag+1, r)
	}

	qSort(points, 0, len(points)-1)

	ans := 0
	for i := 1; i < len(points); i++ {
		ans = tool.Max(ans, points[i][0]-points[i-1][0])
	}

	return ans
}
