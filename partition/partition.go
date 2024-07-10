package partition

import . "github.com/wenxuan7/leetcode/data"

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
