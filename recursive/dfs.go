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

// combinationSum 39. 组合总和
// https://leetcode.cn/problems/combination-sum/description/?envType=study-plan-v2&envId=top-interview-150
func combinationSum(candidates []int, target int) [][]int {
	sum, endCnt, l := 0, 0, len(candidates)
	ans := make([][]int, 0)
	dp := make([]int, 0, l)
	dpCnt := make([]int, 0, l)

	var dfs func(start int)
	dfs = func(start int) {
		if sum > target {
			return
		}
		if sum == target {
			cp := make([]int, 0, endCnt)
			for i, v := range dp {
				for j := 0; j < dpCnt[i]; j++ {
					cp = append(cp, v)
				}
			}
			ans = append(ans, cp)
			return
		}
		if start >= l {
			return
		}

		for i := start; i < l; i++ {
			num := candidates[i]
			if num > target-sum {
				continue
			}
			for cnt := (target - sum) / num; cnt > 0; cnt-- {
				dp = append(dp, num)
				dpCnt = append(dpCnt, cnt)
				sum += num * cnt
				endCnt += cnt
				dfs(i + 1)
				endCnt -= cnt
				sum -= num * cnt
				dpCnt = dpCnt[:len(dpCnt)-1]
				dp = dp[:len(dp)-1]
			}
		}
	}

	dfs(0)
	return ans
}

// generateParenthesis 22. 括号生成
// https://leetcode.cn/problems/generate-parentheses/description/?envType=study-plan-v2&envId=top-interview-150
func generateParenthesis(n int) []string {
	dp := make([]byte, 0, n*2)
	ret := make([]string, 0, n*2-1)

	var dfs func(l, r int)
	dfs = func(l, r int) {
		if l == n && r == n {
			cp := make([]byte, len(dp))
			copy(cp, dp)
			ret = append(ret, string(cp))
			return
		}

		if l < n {
			dp = append(dp, '(')
			dfs(l+1, r)
			dp = dp[:len(dp)-1]
		}
		if r < l {
			dp = append(dp, ')')
			dfs(l, r+1)
			dp = dp[:len(dp)-1]
		}
	}
	dfs(0, 0)
	return ret
}
