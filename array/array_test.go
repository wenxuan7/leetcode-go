package array

import (
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

var (
	threeSumData = [][]int{
		{-1, 0, 1, 2, -1, -4},
	}
	threeSumActual = [][][]int{
		{{-1, -1, 2}, {-1, 0, 1}},
	}
)

func TestThreeSum(t *testing.T) {
	for i := range threeSumData {
		verifySecArr(t, i, threeSum(threeSumData[i]), threeSumActual[i])
	}
}

var (
	removeDuplicatesData = [][]int{
		{1, 1, 1, 1},
		{},
		{1},
		{1, 1, 2, 2, 3, 4, 5, 5, 6},
	}
	removeDuplicatesActual = [][]int{
		{1},
		{},
		{1},
		{1, 2, 3, 4, 5, 6},
	}
)

func TestRemoveDuplicates(t *testing.T) {
	for i := range removeDuplicatesData {
		result := removeDuplicates(removeDuplicatesData[i])
		verifyArr(t, i, removeDuplicatesData[i][:result], removeDuplicatesActual[i])
	}
}

var (
	rotateData = [][]int{
		{},
		{1, 2},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}
	rotateK = []int{
		1,
		-1,
		2,
		7,
	}
	rotateActual = [][]int{
		{},
		{1, 2},
		{4, 5, 1, 2, 3},
		{4, 5, 1, 2, 3},
	}
)

func TestRotate(t *testing.T) {
	for i := range rotateData {
		rotate(rotateData[i], rotateK[i])
		verifyArr(t, i, rotateData[i], rotateActual[i])
	}
}

var (
	mergeData1 = [][]int{
		{},
		{0},
		{1},
		{1, 3, 4, 6, 7, 0, 0, 0},
		{8, 9, 10, 0, 0, 0},
		{1, 3, 4, 0, 0, 0},
	}
	mergeM = []int{
		0,
		0,
		1,
		5,
		3,
		3,
	}
	mergeData2 = [][]int{
		{},
		{1},
		{0},
		{8, 9, 10},
		{1, 2, 3},
		{2, 5, 8},
	}
	mergeN = []int{
		0,
		1,
		0,
		3,
		3,
		3,
	}

	mergeActual = [][]int{
		{},
		{1},
		{1},
		{1, 3, 4, 6, 7, 8, 9, 10},
		{1, 2, 3, 8, 9, 10},
		{1, 2, 3, 4, 5, 8},
	}
)

func TestMerge(t *testing.T) {
	for i := range mergeData1 {
		merge(mergeData1[i], mergeM[i], mergeData2[i], mergeN[i])
		verifyArr(t, i, mergeData1[i], mergeActual[i])
	}
}

var (
	twoSumData = [][]int{
		{},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}
	twoSumTarget = []int{
		1,
		9,
		12,
	}
	twoSumActual = [][]int{
		{},
		{3, 4},
		{},
	}
)

func TestTwoSum(t *testing.T) {
	for i := range twoSumData {
		result := twoSum(twoSumData[i], twoSumTarget[i])
		verifyArr(t, i, result, twoSumActual[i])
	}
}

var (
	plusOneData = [][]int{
		{},
		{1, 2, 3},
		{0},
		{1, 2, 9},
		{9, 9, 9},
	}
	plusOneActual = [][]int{
		{},
		{1, 2, 4},
		{1},
		{1, 3, 0},
		{1, 0, 0, 0},
	}
)

func TestPlusOne(t *testing.T) {
	for i := range plusOneData {
		result := plusOne(plusOneData[i])
		verifyArr(t, i, result, plusOneActual[i])
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
