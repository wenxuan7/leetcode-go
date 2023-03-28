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
