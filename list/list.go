package list

import . "github.com/leetcode/data"

// 2. 两数相加
// addTwoNumbers https://leetcode.cn/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	preH := &ListNode{}
	pre := preH
	multi := 0
	for l1 != nil || l2 != nil {
		sum := multi
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum, multi = sum%10, sum/10
		curr := &ListNode{Val: sum, Next: nil}
		if preH.Next == nil {
			preH.Next = curr
		}
		pre.Next = curr
		pre = curr
	}

	if multi > 0 {
		pre.Next = &ListNode{Val: multi, Next: nil}
	}

	return preH.Next
}
