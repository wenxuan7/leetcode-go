package bs

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
