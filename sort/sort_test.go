package sort

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := [][]int{
		{2, 0, 2, 1, 1, 0},
		{2, 0, 1},
	}
	actual := [][]int{
		{0, 0, 1, 1, 2, 2},
		{0, 1, 2},
	}

	for i := range nums {
		sortColors(nums[i])
		assert.VerifyArr(t, i, nums[i], actual[i])
	}
}

func TestMaxWidthOfVerticalArea(t *testing.T) {
	points := [][][]int{
		{{8, 7}, {9, 9}, {7, 4}, {9, 7}},
		{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}},
	}
	actual := []int{
		1,
		3,
	}

	for i := range points {
		result := maxWidthOfVerticalArea(points[i])
		assert.Verify(t, i, result, actual[i])
	}
}
