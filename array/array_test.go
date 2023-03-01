package array

import (
	"github.com/leetcode-go/assert"
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
		assert.Verify(t, i, maxArea(heightArr[i]), maxAreaArr[i])
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
		assert.VerifyArr(t, i, data[i], actual[i])
	}
}

func TestClimbStairs(t *testing.T) {
	var (
		data   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
		actual = []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	)

	for i := range data {
		assert.Verify(t, i, climbStairs(data[i]), actual[i])
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
		assert.Verify2Arr(t, i, threeSum(data[i]), actual[i])
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
		assert.VerifyArr(t, i, data[i][:result], actual[i])
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
		assert.VerifyArr(t, i, data[i], actual[i])
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
		assert.VerifyArr(t, i, data1[i], actual[i])
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
		assert.VerifyArr(t, i, result, actual[i])
	}
}
