package recursive

// combine 77. 组合
// https://leetcode.cn/problems/combinations/submissions/545581153/?envType=study-plan-v2&envId=top-interview-150
func combine(n int, k int) [][]int {
	if k < 1 {
		return [][]int{}
	}

	var ans [][]int
	dp := make([]int, 0, k)
	var dfs func(int)
	dfs = func(start int) {
		if len(dp)+1+n-start < k {
			return
		}
		if len(dp) == k {
			cp := make([]int, len(dp))
			copy(cp, dp)
			ans = append(ans, cp)
			return
		}

		for i := start; i <= n; i++ {
			dp = append(dp, i)
			dfs(i + 1)
			dp = dp[:len(dp)-1]
		}
	}
	dfs(1)
	return ans
}

// permute 46. 全排列
// https://leetcode.cn/problems/permutations/description/?envType=study-plan-v2&envId=top-interview-150
func permute(nums []int) [][]int {
	dp := make([]int, 0, len(nums))
	ans := make([][]int, 0, len(nums))
	appeared := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(dp) == len(nums) {
			cp := make([]int, len(dp))
			copy(cp, dp)
			ans = append(ans, cp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if appeared[i] {
				continue
			}
			appeared[i] = true
			dp = append(dp, nums[i])
			dfs()
			dp = dp[:len(dp)-1]
			appeared[i] = false
		}
	}

	dfs()
	return ans
}
