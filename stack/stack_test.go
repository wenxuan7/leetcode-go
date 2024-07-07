package stack

import (
	"github.com/wenxuan7/leetcode/require"
	"testing"
)

func TestCalculate(t *testing.T) {
	data := []string{
		"1+(4+5+2)-3",
		"(1+(4+5+2)-3)+(6+8)",
		"0",
	}
	actual := []int{
		9,
		23,
		0,
	}
	for i := range data {
		ret := calculate(data[i])
		require.EqualN(t, i, ret, actual[i])
	}
}
