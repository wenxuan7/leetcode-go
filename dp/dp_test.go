package dp

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	m := []int{
		3,
		3,
	}
	n := []int{
		7,
		2,
	}
	actual := []int{
		28,
		3,
	}

	for i := range m {
		result := uniquePaths(m[i], n[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][][]int{
		{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		{{0, 1}, {0, 0}},
	}
	actual := []int{
		2,
		1,
	}

	for i := range obstacleGrid {
		result := uniquePathsWithObstacles(obstacleGrid[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	text1 := []string{
		"abcde",
		"abc",
		"abc",
		"bsbininm",
	}
	text2 := []string{
		"ace",
		"abc",
		"def",
		"jmjkbkjkv",
	}
	actual := []int{
		3,
		3,
		0,
		1,
	}

	for i := range text1 {
		result := longestCommonSubsequence(text1[i], text2[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestMinimumTotal(t *testing.T) {
	triangle := [][][]int{
		{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
		{{-10}},
		{{-1}, {-2, -3}},
		{{-1}, {2, 3}, {1, -1, -3}},
	}
	actual := []int{
		11,
		-10,
		-4,
		-1,
	}

	for i := range triangle {
		result := minimumTotal(triangle[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestMaxSubArray(t *testing.T) {
	nums := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
	}
	actual := []int{
		6,
		1,
		23,
	}

	for i := range nums {
		result := maxSubArray(nums[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestMaxProduct(t *testing.T) {
	nums := [][]int{
		{2, 3, -2, 4},
		{-2, 0, -1},
		{-2, 3, -4},
	}
	actual := []int{
		6,
		0,
		24,
	}

	for i := range nums {
		result := maxProduct(nums[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestCoinChange(t *testing.T) {
	coins := [][]int{
		{1, 2, 5},
		{2},
		{1},
	}
	amount := []int{
		11,
		3,
		0,
	}
	actual := []int{
		3,
		-1,
		0,
	}

	for i := range coins {
		result := coinChange(coins[i], amount[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestRob(t *testing.T) {
	nums := [][]int{
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
		{2, 1, 1, 2},
	}
	actual := []int{
		4,
		12,
		4,
	}

	for i := range nums {
		result := rob(nums[i])
		assert.Verify(t, i, result, actual[i])
	}
}
