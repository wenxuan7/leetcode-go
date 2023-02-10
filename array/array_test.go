package array

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
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

	for i := range heightArr {
		verify(t, i, maxArea(heightArr[i]), maxAreaArr[i])
	}
}

func TestMoveZeroes(t *testing.T) {

	var (
		data = [][]int{
			{1, 0, 3, 0, 1, 2, 4},
			{0, 1, 0, 3, 12},
			{0},
		}
		actual = [][]int{
			{1, 3, 1, 2, 4, 0, 0},
			{1, 3, 12, 0, 0},
			{0},
		}
	)

	for i := range data {
		moveZeroes(data[i])
		verifyArr(t, i, data[i], actual[i])
	}
}

func TestClimbStairs(t *testing.T) {
	var (
		data   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
		actual = []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	)

	for i := range data {
		verify(t, i, climbStairs(data[i]), actual[i])
	}
}

func TestThreeSum(t *testing.T) {
	var (
		data = [][]int{
			{-1, 0, 1, 2, -1, -4},
		}
		actual = [][][]int{
			{{-1, -1, 2}, {-1, 0, 1}},
		}
	)

	for i := range data {
		verifySecArr(t, i, threeSum(data[i]), actual[i])
	}
}

func TestRemoveDuplicates(t *testing.T) {
	var (
		data = [][]int{
			{1, 1, 1, 1},
			{},
			{1},
			{1, 1, 2, 2, 3, 4, 5, 5, 6},
		}
		actual = [][]int{
			{1},
			{},
			{1},
			{1, 2, 3, 4, 5, 6},
		}
	)

	for i := range data {
		result := removeDuplicates(data[i])
		verifyArr(t, i, data[i][:result], actual[i])
	}
}

func TestRotate(t *testing.T) {
	var (
		data = [][]int{
			{},
			{1, 2},
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4, 5},
		}
		k = []int{
			1,
			-1,
			2,
			7,
		}
		actual = [][]int{
			{},
			{1, 2},
			{4, 5, 1, 2, 3},
			{4, 5, 1, 2, 3},
		}
	)

	for i := range data {
		rotate(data[i], k[i])
		verifyArr(t, i, data[i], actual[i])
	}
}

func TestMerge(t *testing.T) {
	var (
		data1 = [][]int{
			{},
			{0},
			{1},
			{1, 3, 4, 6, 7, 0, 0, 0},
			{8, 9, 10, 0, 0, 0},
			{1, 3, 4, 0, 0, 0},
		}
		m = []int{
			0,
			0,
			1,
			5,
			3,
			3,
		}
		data2 = [][]int{
			{},
			{1},
			{0},
			{8, 9, 10},
			{1, 2, 3},
			{2, 5, 8},
		}
		n = []int{
			0,
			1,
			0,
			3,
			3,
			3,
		}

		actual = [][]int{
			{},
			{1},
			{1},
			{1, 3, 4, 6, 7, 8, 9, 10},
			{1, 2, 3, 8, 9, 10},
			{1, 2, 3, 4, 5, 8},
		}
	)

	for i := range data1 {
		merge(data1[i], m[i], data2[i], n[i])
		verifyArr(t, i, data1[i], actual[i])
	}
}

func TestTwoSum(t *testing.T) {
	var (
		data = [][]int{
			{},
			{1, 2, 3, 4, 5, 6},
			{1, 2, 3, 4, 5, 6},
		}
		target = []int{
			1,
			9,
			12,
		}
		actual = [][]int{
			{},
			{3, 4},
			{},
		}
	)

	for i := range data {
		result := twoSum(data[i], target[i])
		verifyArr(t, i, result, actual[i])
	}
}

func TestPlusOne(t *testing.T) {
	var (
		data = [][]int{
			{},
			{1, 2, 3},
			{0},
			{1, 2, 9},
			{9, 9, 9},
		}
		actual = [][]int{
			{},
			{1, 2, 4},
			{1},
			{1, 3, 0},
			{1, 0, 0, 0},
		}
	)

	for i := range data {
		result := plusOne(data[i])
		verifyArr(t, i, result, actual[i])
	}
}

// verify 校验值是否相等
func verify(t *testing.T, caseIndex int, result int, actual int) {
	if result != actual {
		t.Fatalf("结果与实际值不同, caseIndex: %d, result: %d, actual: %d",
			caseIndex, result, actual)
	}
}

// verifyArr 校验数组全部值是否相等
func verifyArr(t *testing.T, caseIndex int, result []int, actual []int) {
	if len(result) != len(actual) {
		t.Fatalf("len must be equal, result: %v, resultLen: %d, actual: %v, actualLen: %d",
			result, len(result), actual, len(actual))
	}

	for i := range actual {
		if result[i] != actual[i] {
			t.Fatalf("结果与实际值不同, caseIndex: %d, result: %d, actual: %d, i: %d",
				caseIndex, result, actual, i)
		}
	}
}

// verifySecArr 校验二维数组全部值是否相等
func verifySecArr(t *testing.T, caseIndex int, result [][]int, actual [][]int) {
	if len(result) != len(actual) {
		t.Fatalf("len must be equal, result: %v, resultLen: %d, actual: %v, actualLen: %d",
			result, len(result), actual, len(actual))
	}

	for i := range actual {
		for j := range actual[i] {
			if actual[i][j] != result[i][j] {
				t.Fatalf("结果与实际值不同, caseIndex: %d, result: %v, actual: %v, i: %d",
					caseIndex, result, actual, i)
			}
		}
	}
}
