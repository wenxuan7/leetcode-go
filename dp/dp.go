package dp

import (
	"github.com/leetcode-go/tool"
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

// maxSubArray
// 53. 最大子数组和
// https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

// maxProduct
// 152. 乘积最大子数组
// https://leetcode.cn/problems/maximum-product-subarray/
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = tool.Max(mx*nums[i], tool.Max(nums[i], mn*nums[i]))
		minF = tool.Min(mx*nums[i], tool.Min(nums[i], mn*nums[i]))
		ans = tool.Max(maxF, ans)
	}

	return ans
}

// coinChange
// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	if len(coins) == 0 {
		return -1
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1

		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}

			if dp[i-coins[j]]+1 < dp[i] {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

// rob
// 198. 打家劫舍
// https://leetcode.cn/problems/house-robber/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	// 第一个为本身不算进去
	// 第二个为本身算进去
	dp := make([][]int, len(nums))
	dp[0] = []int{0, nums[0]}
	dp[1] = []int{nums[0], nums[1]}

	for i := 2; i < len(nums); i++ {
		dp[i] = make([]int, 2)
		dp[i][0], dp[i][1] = tool.Max(dp[i-1][1], dp[i-1][0]), tool.Max(dp[i-2][1]+nums[i], dp[i-2][0]+nums[i])
	}

	if dp[len(nums)-1][0] > dp[len(nums)-1][1] {
		return dp[len(nums)-1][0]
	} else {
		return dp[len(nums)-1][1]
	}
}

// rob2
// 213. 打家劫舍 II
// https://leetcode.cn/problems/house-robber-ii/description/
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return tool.Max(nums[0], nums[1])
	}

	return tool.Max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

// maxProfit
// 121. 买卖股票的最佳时机
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// dp[i] = max(dp[i-1], prices[i]-min)
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}

	min, dp := prices[0], make([]int, l)
	dp[0] = 0
	for i := 1; i < l; i++ {
		if prices[i] < min {
			min = prices[i]
			dp[i] = dp[i-1]
			continue
		}

		dp[i] = tool.Max(prices[i]-min, dp[i-1])
	}

	return dp[l-1]
}

// maxProfitII
// 122. 买卖股票的最佳时机 II
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfitII(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}

	dp := []int{0, -prices[0]}
	for i := 1; i < l; i++ {
		dp[0], dp[1] = tool.Max(dp[0], dp[1]+prices[i]), tool.Max(dp[1], dp[0]-prices[i])
	}

	return dp[0]
}

// tribonacci
// 1137. 第 N 个泰波那契数
// https://leetcode.cn/problems/n-th-tribonacci-number/
func tribonacci(n int) int {
	if n < 0 {
		return 0
	}

	dp := []int{0, 1, 1}
	if n < 3 {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		dp[0], dp[1], dp[2] = dp[1], dp[2], dp[0]+dp[1]+dp[2]
	}
	return dp[2]
}

// countSubstrings
// 1638. 统计只差一个字符的子串数目
// https://leetcode.cn/problems/count-substrings-that-differ-by-one-character/
func countSubstrings(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}

	ans := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			k := 0
			df := 0
			for ; i+k < len(s) && j+k < len(t); k++ {
				if s[i+k] != t[j+k] {
					df++
				}
				if df == 1 {
					ans++
				}
				if df > 1 {
					break
				}
			}
		}
	}
	return ans
}

// fib
// 509. 斐波那契数
// https://leetcode.cn/problems/fibonacci-number/
func fib(n int) int {
	if n < 1 {
		return 0
	}

	dp := []int{0, 1}
	if n < 2 {
		return dp[n]
	}

	for i := 2; i <= n; i++ {
		dp[0], dp[1] = dp[1], dp[0]+dp[1]
	}
	return dp[1]
}

// countVowelStrings
// 1641. 统计字典序元音字符串的数目
// https://leetcode.cn/problems/count-sorted-vowel-strings/
func countVowelStrings(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 5
	case 2:
		return 15
	}

	dp0, dp1 := []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}
	for i := 3; i <= n; i++ {
		copy(dp0, dp1)

		for j, _ := range dp1 {
			temp := 0
			for k := 0; k <= j; k++ {
				temp += dp0[k]
			}
			dp1[j] = temp
		}
	}

	ans := 0
	for _, num := range dp1 {
		ans += num
	}
	return ans
}
