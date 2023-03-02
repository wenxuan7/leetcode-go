package array

// mySqrt
// 69. x 的平方根
// https://leetcode.cn/problems/sqrtx/
func mySqrt(x int) int {
	l, r, mid, product := 0, 46340, 0, 0

	for l < r {
		mid = (r-l+1)>>1 + l
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
