package array

import (
	"fmt"
	"testing"
)

var (
	heightArr = [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
		{1},
		{1, 2, 3, 1},
		{3, 3, 3, 3},
		{1, 2, 4, 3},
	}
	maxAreaArr = []int{49, 1, 0, 3, 9, 4}
)

func TestMaxArea(t *testing.T) {
	for i := range heightArr {
		verify(t, i, maxArea(heightArr[i]), maxAreaArr[i])
	}
}

func verify(t *testing.T, caseIndex int, result int, actual int) {
	if result != actual {
		t.Fatal(fmt.Sprintf("结果与实际值不同, caseIndex: %d, result: %d, actual: %d", caseIndex, result, actual))
	}
}
