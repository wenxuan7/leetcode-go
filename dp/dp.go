package dp

import (
	"math"
)

// uniquePaths
// 62. 不同路径
// https://leetcode.cn/problems/unique-paths/
func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}

	rows := make([][]int, m)
	for i := 0; i < m; i++ {
		rows[i] = make([]int, n)
		if i == 0 {
			rows[0][0] = 1
		}
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				rows[i][j] += rows[i-1][j]
			}
			if j-1 >= 0 {
				rows[i][j] += rows[i][j-1]
			}
		}
	}

	return rows[m-1][n-1]
}

// uniquePathsWithObstacles
// 63. 不同路径 II
// https://leetcode.cn/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	rows := make([][]int, m)
	for i := 0; i < m; i++ {
		rows[i] = make([]int, n)
		if i == 0 {
			rows[0][0] = 1
		}
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}

			if i-1 >= 0 {
				rows[i][j] += rows[i-1][j]
			}
			if j-1 >= 0 {
				rows[i][j] += rows[i][j-1]
			}
		}
	}

	return rows[m-1][n-1]
}

// longestCommonSubsequence
// 1143. 最长公共子序列
// https://leetcode.cn/problems/longest-common-subsequence/
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	return dp[m][n]
}

// minimumTotal
// 120. 三角形最小路径和
// https://leetcode.cn/problems/triangle/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		nums := triangle[i]
		dp[i] = nums[i] + dp[i-1]
		for j := i - 1; j > 0; j-- {
			if dp[j-1] < dp[j] {
				dp[j] = dp[j-1] + nums[j]
			} else {
				dp[j] = dp[j] + nums[j]
			}
		}

		dp[0] = nums[0] + dp[0]
	}

	min := math.MaxInt
	for _, num := range dp {
		if num < min {
			min = num
		}
	}

	return min
}
