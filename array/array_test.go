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

var (
	moveZeroesBefore = [][]int{
		{1, 0, 3, 0, 1, 2, 4},
		{0, 1, 0, 3, 12},
		{0},
	}
	moveZeroesAfter = [][]int{
		{1, 3, 1, 2, 4, 0, 0},
		{1, 3, 12, 0, 0},
		{0},
	}
)

func TestMoveZeroes(t *testing.T) {
	for i := range moveZeroesBefore {
		moveZeroes(moveZeroesBefore[i])
		verifyArr(t, i, moveZeroesBefore[i], moveZeroesAfter[i])
	}
}

var (
	climbStairsBefore = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	climbStairsAfter  = [...]int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
)

func TestClimbStairs(t *testing.T) {
	for i := range climbStairsBefore {
		verify(t, i, climbStairs(climbStairsBefore[i]), climbStairsAfter[i])
	}
}

func verify(t *testing.T, caseIndex int, result int, actual int) {
	if result != actual {
		t.Fatal(fmt.Sprintf("结果与实际值不同, caseIndex: %d, result: %d, actual: %d",
			caseIndex, result, actual))
	}
}

func verifyArr(t *testing.T, caseIndex int, result []int, actual []int) {
	if len(result) != len(actual) {
		t.Fatal(fmt.Sprintf("len must be equal, resultLen: %d, actualLen: %d",
			len(result), len(actual)))
	}

	for i := range actual {
		if result[i] != actual[i] {
			t.Fatal(fmt.Sprintf("结果与实际值不同, caseIndex: %d, result: %d, actual: %d, i: %d",
				caseIndex, result, actual, i))
		}
	}
}
