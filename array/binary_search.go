package array

// mySqrt
// 69. x 的平方根
// https://leetcode.cn/problems/sqrtx/
func mySqrt(x int) int {
	l, r, mid, product := 0, 46340, 0, 0

	for l < r {
		mid = ((r - l + 1) >> 1) + l
		product = mid * mid
		if product > x {
			r = mid - 1
		} else if product == x {
			return mid
		} else {
			l = mid
		}
	}
	return l
}

// isPerfectSquare
// 367. 有效的完全平方数
// https://leetcode.cn/problems/valid-perfect-square/
func isPerfectSquare(num int) bool {
	l, r, mid, product := 0, 46340, 0, 0

	for l <= r {
		mid = ((r - l) >> 1) + l
		product = mid * mid
		if product > num {
			r = mid - 1
		} else if product == num {
			return true
		} else {
			l = mid + 1
		}
	}
	return false
}

// search
// 33. 搜索旋转排序数组
// https://leetcode.cn/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r, mid := 0, len(nums)-1, 0
	for l <= r {
		mid = ((r - l) >> 1) + l
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			if target >= nums[mid] || nums[l] > target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else if nums[mid] < nums[r] {
			if target > nums[r] || nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}

// searchMatrix
// 74. 搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	l, r, mid := 0, len(matrix)-1, 0
	for l < r {
		mid = ((r - l + 1) >> 1) + l
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}

	row := l
	l, r = 0, len(matrix[row])-1
	for l <= r {
		mid = ((r - l + 1) >> 1) + l
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
