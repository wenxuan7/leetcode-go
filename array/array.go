package array

import "sort"

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

// 15. 三数之和
// https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	if nums == nil {
		return [][]int{}
	}

	l := len(nums)
	if l < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	if nums[l-1] < 0 {
		return [][]int{}
	}

	var ans [][]int
	for i := range nums {
		if nums[i] > 0 {
			return ans
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, l-1; j < k; {
			target := nums[j] + nums[k] + nums[i]

			if target > 0 {
				k--
			} else if target < 0 {
				j++
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--

				for j < k && nums[j] == nums[j-1] {
					j++
				}

				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return ans
}

// removeDuplicates
// 26. 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) < 2 {
		return 1
	}

	fast, slow := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow-1] {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}

	return slow
}

// rotate
// 189. 轮转数组
// https://leetcode.cn/problems/rotate-array/
func rotate(nums []int, k int) {
	if k < 1 || nums == nil ||
		len(nums) == 0 {
		return
	}

	k %= len(nums)
	rotateAll(nums)
	rotateAll(nums[0:k])
	rotateAll(nums[k:])
}

func rotateAll(nums []int) {
	l, r := 0, len(nums)-1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
