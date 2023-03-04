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
