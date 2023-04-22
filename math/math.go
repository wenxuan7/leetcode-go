package math

// smallestEvenMultiple
// 2413. 最小偶倍数
// https://leetcode.cn/problems/smallest-even-multiple/
func smallestEvenMultiple(n int) int {
	if n <= 0 {
		return 0
	}

	if n&1 == 1 {
		return n * 2
	}
	return n
}
