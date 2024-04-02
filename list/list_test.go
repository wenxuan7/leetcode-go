package list

import (
	. "github.com/leetcode/data"
	"github.com/leetcode/require"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	data1 := []*ListNode{
		require.HListNodeFromArray(4, 3, 5),
		require.HListNodeFromArray(5, 5, 5),
		nil,
	}
	data2 := []*ListNode{
		require.HListNodeFromArray(5, 8, 7),
		require.HListNodeFromArray(8, 8),
		nil,
	}
	actual := []*ListNode{
		require.HListNodeFromArray(9, 1, 3, 1),
		require.HListNodeFromArray(3, 4, 6),
		nil,
	}

	for i := range actual {
		result := addTwoNumbers(data1[i], data2[i])
		require.EqualListNode(t, i, result, actual[i])
	}
}
