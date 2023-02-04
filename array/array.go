package array

// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/
func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}

	l, r, max := 0, len(height)-1, 0
	for l < r {
		wid, area := r-l, 0
		if height[r] > height[l] {
			area = height[l] * wid
			oldL := l
			for l < r && height[l] <= height[oldL] {
				l++
			}
		} else {
			area = height[r] * wid
			oldR := r
			for l < r && height[r] <= height[oldR] {
				r--
			}
		}

		if area > max {
			max = area
		}
	}

	return max
}

// 283. 移动零
// https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

// 70. 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/
func climbStairs(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	dp := [...]int{1, 2}
	for i := 3; i <= n; i++ {
		dp[0], dp[1] = dp[1], dp[0]+dp[1]
	}

	return dp[1]
}
