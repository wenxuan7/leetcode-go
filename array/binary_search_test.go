package array

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestMySqrt(t *testing.T) {
	x := []int{
		4,
		8,
	}
	actual := []int{
		2,
		2,
	}

	for i := range x {
		result := mySqrt(x[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestIsPerfectSquare(t *testing.T) {
	num := []int{
		16,
		14,
		9,
	}
	actual := []bool{
		true,
		false,
		true,
	}

	for i := range num {
		result := isPerfectSquare(num[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestSearch(t *testing.T) {
	num := [][]int{
		{4, 5, 6, 7, 0, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{1},
		{1},
		{3, 1},
	}
	target := []int{
		0,
		3,
		0,
		1,
		1,
	}
	actual := []int{
		4,
		-1,
		-1,
		0,
		1,
	}

	for i := range num {
		result := search(num[i], target[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestSearchMatrix(t *testing.T) {
	matrix := [][][]int{
		{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		{{1}},
	}
	target := []int{
		3,
		13,
		1,
	}
	actual := []bool{
		true,
		false,
		true,
	}

	for i := range matrix {
		result := searchMatrix(matrix[i], target[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestFindMin(t *testing.T) {
	nums := [][]int{
		{3, 4, 5, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{11, 13, 15, 17},
	}
	actual := []int{
		1,
		0,
		11,
	}

	for i := range nums {
		result := findMin(nums[i])
		assert.Verify(t, i, result, actual[i])
	}
}
