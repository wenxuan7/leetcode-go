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

func TestQSort(t *testing.T) {
	nums := [][]int{
		{6, 5, 4, 3, 2, 1},
		{7, 8, 1, 5, 2, 3},
		{1, 1, 1, 2, 2, 3, 4, 5},
	}
	actual := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 5, 7, 8},
		{1, 1, 1, 2, 2, 3, 4, 5},
	}

	for i := range nums {
		qSort(nums[i], 0, len(nums[i])-1)
		assert.VerifyArr(t, i, nums[i], actual[i])
	}
}
