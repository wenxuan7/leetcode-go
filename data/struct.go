package data

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 二叉树节点
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Node n叉树节点
type Node struct {
	Val      int
	Children []*Node
}
