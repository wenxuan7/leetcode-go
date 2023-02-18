package queue

import . "leetcode-go/data"

// 102. 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		level := make([]int, 0)

		levelCount := len(q)
		for i := 0; i < levelCount; i++ {
			level = append(level, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		q = q[levelCount:]
		ans = append(ans, level)
	}

	return ans
}
