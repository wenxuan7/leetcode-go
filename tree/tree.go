package tree

import . "leetcode-go/data"

// inorderTraversal
// 94. 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/
var inorder []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	inorder = make([]int, 0)
	recursiveInorder(root)
	return inorder
}

func recursiveInorder(tn *TreeNode) {
	if tn == nil {
		return
	}

	recursiveInorder(tn.Left)
	inorder = append(inorder, tn.Val)
	recursiveInorder(tn.Right)
}

// preorderTraversal
// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/
var pre []int

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	pre = make([]int, 0)
	recursivePreorder(root)
	return pre
}

func recursivePreorder(tn *TreeNode) {
	if tn == nil {
		return
	}

	pre = append(pre, tn.Val)
	recursivePreorder(tn.Left)
	recursivePreorder(tn.Right)
}

// postorder
// 590. N 叉树的后序遍历
// https://leetcode.cn/problems/n-ary-tree-postorder-traversal/
var post []int

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	post = make([]int, 0)
	recursivePostorder(root)
	return post
}

func recursivePostorder(root *Node) {
	if root == nil {
		return
	}

	for _, child := range root.Children {
		recursivePostorder(child)
	}
	post = append(post, root.Val)
}

// 589. N 叉树的前序遍历
// https://leetcode.cn/problems/n-ary-tree-preorder-traversal/
func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	recursivePre(root, &ans)
	return ans
}

func recursivePre(root *Node, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)
	for _, child := range root.Children {
		recursivePre(child, ans)
	}
}

// levelOrder
// 429. N 叉树的层序遍历
// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)
	recursiveLevelOrder(root, 0, &ans)
	return ans
}

func recursiveLevelOrder(root *Node, level int, ans *[][]int) {
	if root == nil {
		return
	}

	if level >= len(*ans) {
		nums := []int{root.Val}
		*ans = append(*ans, nums)
	} else {
		(*ans)[level] = append((*ans)[level], root.Val)
	}

	for _, child := range root.Children {
		recursiveLevelOrder(child, level+1, ans)
	}
}
