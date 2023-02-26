package greedy

import "sort"

// lemonadeChange
// 860. 柠檬水找零
// https://leetcode.cn/problems/lemonade-change/
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five <= 0 {
				return false
			}
			five--
			ten++
		} else {
			if five <= 0 {
				return false
			}

			if ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}

// maxProfit
// 122. 买卖股票的最佳时机 II
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	ans := 0
	for i := 1; i < len(prices); i++ {
		sub := prices[i] - prices[i-1]
		if sub > 0 {
			ans += sub
		}
	}

	return ans
}

// findContentChildren
// 455. 分发饼干
// https://leetcode.cn/problems/assign-cookies/
func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}

	if len(g) == 0 {
		return len(s)
	}

	sort.Ints(g)
	sort.Ints(s)

	ans := 0
	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if s[j] >= g[i] {
			ans++
			i++
		}
	}

	return ans
}
