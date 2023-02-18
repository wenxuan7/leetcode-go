package queue

import (
	. "leetcode-go/data"
)

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

// minMutation
// 433. 最小基因变化
// https://leetcode.cn/problems/minimum-genetic-mutation/
func minMutation(startGene string, endGene string, bank []string) int {
	noExist := 0
	for i := 0; i < len(bank); i++ {
		if bank[i] != endGene {
			noExist++
		}
	}
	if noExist == len(bank) {
		return -1
	}

	q := make([]string, 0)
	q = append(q, startGene)

	ans := 1
	for len(q) > 0 {
		// 防止死循环
		if ans > len(bank) {
			return -1
		}
		levelCount := len(q)
		for i := 0; i < levelCount; i++ {
			if isModify(q[i], endGene) {
				return ans
			}

			for j := 0; j < len(bank); j++ {
				if isModify(q[i], bank[j]) {
					q = append(q, bank[j])
				}
			}
		}

		if len(q) == levelCount {
			return -1
		}
		q = q[levelCount:]
		ans++
	}

	return -1
}

func isModify(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	bs1, bs2 := []byte(str1), []byte(str2)
	count := 0
	for i := 0; i < len(bs1); i++ {
		if bs1[i] != bs2[i] {
			count++
		}

		if count > 1 {
			return false
		}
	}

	return count == 1
}
