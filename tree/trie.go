package tree

import (
	. "github.com/wenxuan7/leetcode/data"
	"math"
)

var maxSum int

// maxPathSum 124. 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/?envType=study-plan-v2&envId=top-interview-150
func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt
	maxGain(root)
	return maxSum
}

func maxGain(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := max(maxGain(root.Left), 0), max(maxGain(root.Right), 0)
	sum := root.Val + l + r
	maxSum = max(maxSum, sum)
	return root.Val + max(l, r)
}
