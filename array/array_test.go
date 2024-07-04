package array

import (
	"github.com/leetcode/require"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	data1 := [][]int{
		{1, 2},
		{0},
		{3, 4},
	}
	data2 := [][]int{
		{3, 4},
		{1, 2},
		{5},
	}

	actual := []float64{
		2.5,
		1,
		4,
	}

	for i := range actual {
		result := findMedianSortedArrays(data1[i], data2[i])
		EqualFloat64 := require.EqualNumber[float64]
		EqualFloat64(t, i, result, actual[i])
	}
}
