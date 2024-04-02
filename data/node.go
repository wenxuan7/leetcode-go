package data

import (
	"strconv"
	"strings"
)

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "<nil>"
	}

	bd := strings.Builder{}
	bd.WriteString("[")
	curr := l
	for curr != nil {
		bd.WriteString(strconv.Itoa(curr.Val))
		if curr.Next != nil {
			bd.WriteString(", ")
		}
		curr = curr.Next
	}
	bd.WriteString("]")

	return bd.String()
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
