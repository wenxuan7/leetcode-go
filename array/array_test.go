package array

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := [][]int{
		{1, 3},
		{1, 2},
	}
	nums2 := [][]int{
		{2},
		{3, 4},
	}
	actual := []float64{
		2,
		2.5,
	}
	for i := range nums1 {
		assert.Verify(t, i, findMedianSortedArrays(nums1[i], nums2[i]), actual[i])
	}
}
