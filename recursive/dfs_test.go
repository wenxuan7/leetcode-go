package recursive

import (
	"fmt"
	"github.com/wenxuan7/leetcode/require"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	data := [][]int{
		{2, 3, 6, 7},
	}
	dataTarget := []int{
		7,
	}

	for i := range data {
		sum := combinationSum(data[i], dataTarget[i])
		fmt.Println(sum)
	}
}

func TestGenerateParenthesis(t *testing.T) {
	data := []int{
		3,
	}
	actual := [][]string{
		{"((()))", "(()())", "(())()", "()(())", "()()()"},
	}
	for i := range data {
		ret := generateParenthesis(data[i])
		require.Equal1D(t, i, actual[i], ret)

	}
}
