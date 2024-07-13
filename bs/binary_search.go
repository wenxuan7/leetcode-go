package bs

import "math"

// searchInsert 35. 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-interview-150
func searchInsert(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l < r {
		mid = l + (r-l)>>1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// searchMatrix 74. 搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/submissions/546108803/?envType=study-plan-v2&envId=top-interview-150
func searchMatrix(matrix [][]int, target int) bool {
	l, r, mid := 0, len(matrix)-1, 0
	for l < r {
		mid = l + (r-l+1)>>1
		if target == matrix[mid][0] {
			return true
		} else if target > matrix[mid][0] {
			l = mid
		} else {
			r = mid - 1
		}
	}

	row := l
	l, r = 0, len(matrix[row])-1
	for l <= r {
		mid = l + (r-l)>>1
		if target == matrix[row][mid] {
			return true
		} else if target > matrix[row][mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

// findPeakElement 162. 寻找峰值
// https://leetcode.cn/problems/find-peak-element/description/?envType=study-plan-v2&envId=top-interview-150
func findPeakElement(nums []int) int {
	n := len(nums)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt
		}
		return nums[i]
	}

	l, mid, r := 0, 0, n-1
	mid1N, midN, midN1 := 0, 0, 0
	for {
		mid = l + (r-l)>>1
		mid1N, midN, midN1 = get(mid-1), get(mid), get(mid+1)
		if mid1N < midN && midN > midN1 {
			return mid
		}
		if midN < midN1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}
