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
